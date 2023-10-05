#!/bin/bash

# imports  
. scripts/envVar.sh
. scripts/utils.sh

DELAY="$2"
MAX_RETRY="$3"
VERBOSE="$4"
: ${CHANNEL_NAME:="vehicles"}
: ${DELAY:="3"}
: ${MAX_RETRY:="5"}
: ${VERBOSE:="false"}

: ${CONTAINER_CLI:="docker"}
: ${CONTAINER_CLI_COMPOSE:="${CONTAINER_CLI}-compose"}
infoln "Using ${CONTAINER_CLI} and ${CONTAINER_CLI_COMPOSE}"

if [ ! -d "channel-artifacts" ]; then
	mkdir channel-artifacts
fi

echo "ORDERER_CA :  $ORDERER_CA"
echo "ORDERER_ADMIN_TLS_SIGN_CERT : $ORDERER_ADMIN_TLS_SIGN_CERT"
echo "ORDERER_ADMIN_TLS_PRIVATE_KEY : $ORDERER_ADMIN_TLS_PRIVATE_KEY"

which configtxgen
if [ "$?" -ne 0 ]; then
    fatalln "configtxgen tool not found."
fi

makeGenesisBlock(){
  CHANNEL_NAME="vehicles"
  FABRIC_CFG_PATH=${PWD}/configtx
  # set -x
  configtxgen -profile people -outputBlock ./channel-artifacts/${CHANNEL_NAME}.block -channelID $CHANNEL_NAME
  FABRIC_CFG_PATH=$PWD/../config/
  osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:7053 --ca-file "$ORDERER_CA" --client-cert "$ORDERER_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER_ADMIN_TLS_PRIVATE_KEY" >&log.txt
  cat log.txt
}

#joinChannel 채널이름 Org이름
joinChannel() {
  FABRIC_CFG_PATH=$PWD/../config/
  CHANNEL_NAME=$1
  ORG=$2
  setGlobals $ORG
  BLOCKFILE="./channel-artifacts/${CHANNEL_NAME}.block"
	local rc=1
	local COUNTER=1
	## Sometimes Join takes time, hence retry
	while [ $rc -ne 0 -a $COUNTER -lt $MAX_RETRY ] ; do
    sleep $DELAY
    set -x
    peer channel join -b $BLOCKFILE >&log.txt
    res=$?
    { set +x; } 2>/dev/null
		let rc=$res
		COUNTER=$(expr $COUNTER + 1)
	done
	cat log.txt
	verifyResult $res "After $MAX_RETRY attempts, peer0.${ORG} has failed to join channel '$CHANNEL_NAME' "
}
#joinChannel 채널이름 Org이름
joinChannel2() {
  FABRIC_CFG_PATH=$PWD/../config/
  CHANNEL_NAME=$1
  ORG=$2
  setGlobals2 $ORG
  BLOCKFILE="./channel-artifacts/${CHANNEL_NAME}.block"
	local rc=1
	local COUNTER=1
	## Sometimes Join takes time, hence retry
	while [ $rc -ne 0 -a $COUNTER -lt $MAX_RETRY ] ; do
    sleep $DELAY
    set -x
    peer channel join -b $BLOCKFILE >&log.txt
    res=$?
    { set +x; } 2>/dev/null
		let rc=$res
		COUNTER=$(expr $COUNTER + 1)
	done
	cat log.txt
	verifyResult $res "After $MAX_RETRY attempts, peer0.${ORG} has failed to join channel '$CHANNEL_NAME' "
}
#setAnchorPeer 채널이름 Org이름
setAnchorPeer() {
  CHANNEL_NAME=$1
  ORG=$2
  # echo "CONTAINER_CLI : $CONTAINER_CLI"
  # echo "ORG : $ORG"
  ${CONTAINER_CLI} exec cli ./scripts/setAnchorPeer.sh $ORG $CHANNEL_NAME 
}

makeGenesisBlock

FABRIC_CFG_PATH=$PWD/../config/
infoln "seller를 vehicles에 가입중.."
joinChannel vehicles seller
joinChannel2 vehicles seller
infoln "buyer를 vehicles에 가입중.."
joinChannel vehicles buyer
joinChannel2 vehicles buyer
infoln "inspector를 vehicles에 가입중.."
joinChannel vehicles inspector
joinChannel2 vehicles inspector

infoln "seller를 AnchorPeer 설정 중.."
setAnchorPeer vehicles seller
infoln "buyer를 AnchorPeer 설정 중.."
setAnchorPeer vehicles buyer
infoln "inspector를 AnchorPeer 설정 중.."
setAnchorPeer vehicles inspector

successln "완료!"
