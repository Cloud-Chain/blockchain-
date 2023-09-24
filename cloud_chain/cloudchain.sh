#!/bin/bash

ROOTDIR=$(cd "$(dirname "$0")" && pwd)
export PATH=${ROOTDIR}/../bin:${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=${PWD}/configtx
export VERBOSE=false

pushd ${ROOTDIR} > /dev/null
trap "popd > /dev/null" EXIT

. scripts/utils.sh

: ${CONTAINER_CLI:="docker"}
: ${CONTAINER_CLI_COMPOSE:="${CONTAINER_CLI}-compose"}
infoln "Using ${CONTAINER_CLI} and ${CONTAINER_CLI_COMPOSE}"

SOCK="${DOCKER_HOST:-/var/run/docker.sock}"
DOCKER_SOCK="${SOCK##unix://}"

export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
CRYPTO="cryptogen"
MAX_RETRY=5
CLI_DELAY=3
CHANNEL_NAME="vehicles"
CC_NAME="NA"
CC_SRC_PATH="NA"
CC_END_POLICY="NA"
CC_COLL_CONFIG="NA"
CC_INIT_FCN="NA"
COMPOSE_FILE_BASE=compose-test-net.yaml
COMPOSE_FILE_COUCH=compose-couch.yaml
COMPOSE_FILE_CA=compose-ca.yaml
CC_SRC_LANGUAGE="NA"
CCAAS_DOCKER_RUN=true
CC_VERSION="1.0"
CC_SEQUENCE=1
DATABASE="leveldb"

function createOrgs() {
  if [ -d "organizations/peerOrganizations" ]; then
    rm -Rf organizations/peerOrganizations && rm -Rf organizations/ordererOrganizations
  fi
  # Create crypto material using cryptogen
  if [ "$CRYPTO" == "cryptogen" ]; then
    which cryptogen
    if [ "$?" -ne 0 ]; then
      fatalln "Cryptogen 오류"
    fi

    infoln "Cryptogen을 이용해서 암호화 파일을 생성"

    infoln "사용자(판매자)의 암호화 파일 생성"
    set -x
    cryptogen generate --config=./organizations/cryptogen/crypto-config-seller.yaml --output="organizations"
    res=$?
    { set +x; } 2>/dev/null
    if [ $res -ne 0 ]; then
      fatalln "Failed to generate certificates..."
    fi

    infoln "사용자(구매자)의 암호화 파일 생성"
    set -x
    cryptogen generate --config=./organizations/cryptogen/crypto-config-buyer.yaml --output="organizations"
    res=$?
    { set +x; } 2>/dev/null
    if [ $res -ne 0 ]; then
      fatalln "Failed to generate certificates..."
    fi
    
    # 검수기관 조직 암호화
    infoln "검수기관 암호화 파일 생성"
    set -x
    cryptogen generate --config=./organizations/cryptogen/crypto-config-inspector.yaml --output="organizations"
    res=$?
    { set +x; } 2>/dev/null
    if [ $res -ne 0 ]; then
      fatalln "Failed to generate certificates..."
    fi

    infoln "오더러 암호화 파일 생성"
    set -x
    cryptogen generate --config=./organizations/cryptogen/crypto-config-orderer.yaml --output="organizations"
    res=$?
    { set +x; } 2>/dev/null
    if [ $res -ne 0 ]; then
      fatalln "Failed to generate certificates..."
    fi

  fi

  # Create crypto material using Fabric CA
  infoln $CRYPTO
  if [ "$CRYPTO" == "Certificate Authorities" ]; then
    infoln "Generating certificates using Fabric CA"
    ${CONTAINER_CLI_COMPOSE} -f compose/$COMPOSE_FILE_CA -f compose/$CONTAINER_CLI/${CONTAINER_CLI}-$COMPOSE_FILE_CA up -d 2>&1

    . organizations/fabric-ca/registerEnroll.sh

    while :
    do
      if [ ! -f "organizations/fabric-ca/seller/ca-cert.pem" ]; then
        sleep 1
      else
        break
      fi
    done

    infoln "Creating Seller Identities"

    createSeller

    infoln "Creating Buyer Identities"

    createBuyer

    infoln "Creating Inspector Identities"

    createInspector

    infoln "Creating Orderer Org Identities"

    createOrderer

  fi

  infoln "Generating CCP files for seller, buyer and inspector"
  ./organizations/ccp-generate.sh
}

function checkPrereqs() {
  ## Check if your have cloned the peer binaries and configuration files.
  peer version > /dev/null 2>&1

  if [[ $? -ne 0 || ! -d "../config" ]]; then
    errorln "Peer binary and configuration files not found.."
    errorln
    errorln "Follow the instructions in the Fabric docs to install the Fabric Binaries:"
    errorln "https://hyperledger-fabric.readthedocs.io/en/latest/install.html"
    exit 1
  fi
  # use the fabric tools container to see if the samples and binaries match your
  # docker images
  LOCAL_VERSION=$(peer version | sed -ne 's/^ Version: //p')
  DOCKER_IMAGE_VERSION=$(${CONTAINER_CLI} run --rm hyperledger/fabric-tools:latest peer version | sed -ne 's/^ Version: //p')

  infoln "LOCAL_VERSION=$LOCAL_VERSION"
  infoln "DOCKER_IMAGE_VERSION=$DOCKER_IMAGE_VERSION"

  if [ "$LOCAL_VERSION" != "$DOCKER_IMAGE_VERSION" ]; then
    warnln "Local fabric binaries and docker images are out of  sync. This may cause problems."
  fi

  for UNSUPPORTED_VERSION in $NONWORKING_VERSIONS; do
    infoln "$LOCAL_VERSION" | grep -q $UNSUPPORTED_VERSION
    if [ $? -eq 0 ]; then
      fatalln "Local Fabric binary version of $LOCAL_VERSION does not match the versions supported by the test network."
    fi

    infoln "$DOCKER_IMAGE_VERSION" | grep -q $UNSUPPORTED_VERSION
    if [ $? -eq 0 ]; then
      fatalln "Fabric Docker image version of $DOCKER_IMAGE_VERSION does not match the versions supported by the test network."
    fi
  done

  ## Check for fabric-ca
  if [ "$CRYPTO" == "Certificate Authorities" ]; then

    fabric-ca-client version > /dev/null 2>&1
    if [[ $? -ne 0 ]]; then
      errorln "fabric-ca-client binary not found.."
      errorln
      errorln "Follow the instructions in the Fabric docs to install the Fabric Binaries:"
      errorln "https://hyperledger-fabric.readthedocs.io/en/latest/install.html"
      exit 1
    fi
    CA_LOCAL_VERSION=$(fabric-ca-client version | sed -ne 's/ Version: //p')
    CA_DOCKER_IMAGE_VERSION=$(docker run --rm hyperledger/fabric-ca:1.5.7 fabric-ca-client version | sed -ne 's/ Version: //p' | head -1)
    infoln "CA_LOCAL_VERSION=$CA_LOCAL_VERSION"
    infoln "CA_DOCKER_IMAGE_VERSION=$CA_DOCKER_IMAGE_VERSION"

    if [ "$CA_LOCAL_VERSION" != "$CA_DOCKER_IMAGE_VERSION" ]; then
      warnln "Local fabric-ca binaries and docker images are out of sync. This may cause problems."
    fi
  fi
}

function networkUp() {
  checkPrereqs
  # generate artifacts if they don't exist
  if [ ! -d "organizations/peerOrganizations" ]; then
    createOrgs
  fi
  COMPOSE_FILES="-f compose/${COMPOSE_FILE_BASE} -f compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_BASE}"
  if [ "${DATABASE}" == "couchdb" ]; then
    infoln "Couchdb 사용"
    COMPOSE_FILES="${COMPOSE_FILES} -f compose/${COMPOSE_FILE_COUCH} -f compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_COUCH}"
  fi
  echo "${DOCKER_SOCK} & ${CONTAINER_CLI_COMPOSE} & ${COMPOSE_FILES}"
  DOCKER_SOCK="${DOCKER_SOCK}" ${CONTAINER_CLI_COMPOSE} ${COMPOSE_FILES} up -d 2>&1
  $CONTAINER_CLI ps -a
  if [ $? -ne 0 ]; then
    fatalln "Unable to start network"
  fi
}

function networkDown() {

  COMPOSE_BASE_FILES="-f compose/${COMPOSE_FILE_BASE} -f compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_BASE}"
  COMPOSE_COUCH_FILES="-f compose/${COMPOSE_FILE_COUCH} -f compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_COUCH}"
  COMPOSE_CA_FILES="-f compose/${COMPOSE_FILE_CA} -f compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_CA}"
  COMPOSE_FILES="${COMPOSE_BASE_FILES} "

  if [ "${CONTAINER_CLI}" == "docker" ]; then
    set -x
    DOCKER_SOCK=$DOCKER_SOCK ${CONTAINER_CLI_COMPOSE} ${COMPOSE_FILES} down --volumes --remove-orphans
  else
    fatalln "Container CLI  ${CONTAINER_CLI} not supported"
  fi
  sudo rm -rf organizations/fabric-ca/seller/msp organizations/fabric-ca/seller/tls-cert.pem organizations/fabric-ca/seller/ca-cert.pem organizations/fabric-ca/seller/IssuerPublicKey organizations/fabric-ca/seller/IssuerRevocationPublicKey organizations/fabric-ca/seller/fabric-ca-server.db
  sudo rm -rf organizations/fabric-ca/buyer/msp organizations/fabric-ca/buyer/tls-cert.pem organizations/fabric-ca/buyer/ca-cert.pem organizations/fabric-ca/buyer/IssuerPublicKey organizations/fabric-ca/buyer/IssuerRevocationPublicKey organizations/fabric-ca/buyer/fabric-ca-server.db
  sudo rm -rf organizations/fabric-ca/ordererOrg/msp organizations/fabric-ca/ordererOrg/tls-cert.pem organizations/fabric-ca/ordererOrg/ca-cert.pem organizations/fabric-ca/ordererOrg/IssuerPublicKey organizations/fabric-ca/ordererOrg/IssuerRevocationPublicKey organizations/fabric-ca/ordererOrg/fabric-ca-server.db
  sudo rm -rf organizations/fabric-ca/inspector/msp organizations/fabric-ca/inspector/tls-cert.pem organizations/fabric-ca/inspector/ca-cert.pem organizations/fabric-ca/inspector/IssuerPublicKey organizations/fabric-ca/inspector/IssuerRevocationPublicKey organizations/fabric-ca/inspector/fabric-ca-server.db
  rm -rf organizations/ordererOrganizations
  rm -rf organizations/peerOrganizations
  rm -rf channel-artifacts
  rm -rf package
}

function setnetwork(){
  bringUpNetwork="false"
  CONTAINERS=($($CONTAINER_CLI ps | grep hyperledger/ | awk '{print $2}'))
  scripts/setChannel.sh $CHANNEL_NAME $CLI_DELAY $MAX_RETRY $VERBOSE

}
## Parse mode
if [[ $# -lt 1 ]] ; then
  printHelp
  exit 0
else
  MODE=$1
  shift
fi


while [[ $# -ge 1 ]] ; do
  key="$1"
  case $key in
  -h )
    printHelp $MODE
    exit 0
    ;;
  -c )
    CHANNEL_NAME="$2"
    shift
    ;;
  -ca )
    CRYPTO="Certificate Authorities"
    ;;
  -r )
    MAX_RETRY="$2"
    shift
    ;;
  -d )
    CLI_DELAY="$2"
    shift
    ;;
  -s )
    DATABASE="$2"
    shift
    ;;
  -ccl )
    CC_SRC_LANGUAGE="$2"
    shift
    ;;
  -ccn )
    CC_NAME="$2"
    shift
    ;;
  -ccv )
    CC_VERSION="$2"
    shift
    ;;
  -ccs )
    CC_SEQUENCE="$2"
    shift
    ;;
  -ccp )
    CC_SRC_PATH="$2"
    shift
    ;;
  -ccep )
    CC_END_POLICY="$2"
    shift
    ;;
  -cccg )
    CC_COLL_CONFIG="$2"
    shift
    ;;
  -cci )
    CC_INIT_FCN="$2"
    shift
    ;;
  -ccaasdocker )
    CCAAS_DOCKER_RUN="$2"
    shift
    ;;
  -verbose )
    VERBOSE=true
    ;;
  * )
    errorln "Unknown flag: $key"
    printHelp
    exit 1
    ;;
  esac
  shift
done

# Are we generating crypto material with this command?
if [ ! -d "organizations/peerOrganizations" ]; then
  CRYPTO_MODE="with crypto from '${CRYPTO}'"
else
  CRYPTO_MODE=""
fi



# Determine mode of operation and printing out what we asked for
if [ "$MODE" == "up" ]; then
  infoln "Starting nodes with CLI timeout of '${MAX_RETRY}' tries and CLI delay of '${CLI_DELAY}' seconds and using database '${DATABASE}' ${CRYPTO_MODE}"
  networkUp
elif [ "$MODE" == "crypto" ]; then
  infoln "암호화 실행"
  http
elif [ "$MODE" == "down" ]; then
  infoln "네트워크 내리고 도커 이미지까지 제거"
  networkDown
elif [ "$MODE" == "set" ]; then
  infoln "네트워크 셋팅 작업"
  setnetwork
elif [ "$MODE" == "install" ]; then
  infoln "체인코드 패키징과 설치 작업"
  . scripts/installChaincode.sh
elif [ "$MODE" == "commit" ]; then
  infoln "체인코드 승인과 커밋 작업"
  . CC_commit.sh
else
  printHelp
  exit 1
fi


