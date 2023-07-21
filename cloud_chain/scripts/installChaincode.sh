#!/bin/bash

source scripts/utils.sh

CHANNEL_NAME=${1:-"mychannel"}
CC_SRC_LANGUAGE=
CC_VERSION=${5:-"1.0"}
CC_SEQUENCE=${6:-"1"}
CC_INIT_FCN=${7:-"NA"}
CC_END_POLICY=${8:-"NA"}
CC_COLL_CONFIG=${9:-"NA"}
DELAY=${10:-"3"}
MAX_RETRY=${11:-"5"}
VERBOSE=${12:-"false"}

println "executing with the following"
println "- CHANNEL_NAME: ${C_GREEN}${CHANNEL_NAME}${C_RESET}"
println "- CC_NAME: ${C_GREEN}${CC_NAME}${C_RESET}"
println "- CC_SRC_PATH: ${C_GREEN}${CC_SRC_PATH}${C_RESET}"
println "- CC_SRC_LANGUAGE: ${C_GREEN}${CC_SRC_LANGUAGE}${C_RESET}"
println "- CC_VERSION: ${C_GREEN}${CC_VERSION}${C_RESET}"
println "- CC_SEQUENCE: ${C_GREEN}${CC_SEQUENCE}${C_RESET}"
println "- CC_END_POLICY: ${C_GREEN}${CC_END_POLICY}${C_RESET}"
println "- CC_COLL_CONFIG: ${C_GREEN}${CC_COLL_CONFIG}${C_RESET}"
println "- CC_INIT_FCN: ${C_GREEN}${CC_INIT_FCN}${C_RESET}"
println "- DELAY: ${C_GREEN}${DELAY}${C_RESET}"
println "- MAX_RETRY: ${C_GREEN}${MAX_RETRY}${C_RESET}"
println "- VERBOSE: ${C_GREEN}${VERBOSE}${C_RESET}"

export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
SOCK="${DOCKER_HOST:-/var/run/docker.sock}"
DOCKER_SOCK="${SOCK##unix://}"

infoln "체인코드 패키징"

# import utils
. scripts/envVar.sh
. scripts/ccutils.sh

packaging(){
    # CC_MARKET_PATH="${PWD}/chaincode/chaincode_market/"
    # CC_REVENUE_PATH="${PWD}/chaincode/chaincode_revenue/"
    # CC_PEOPLE_PATH="${PWD}/chaincode/chaincode_people/"
    CC_TRANSACTION_PATH="${PWD}/chaincode/chaincode_transaction"
    # echo $CC_MARKET_PATH
    # echo $CC_REVENUE_PATH
    # echo $CC_PEOPLE_PATH
    echo $CC_USED_CAR_PATH
    if [ ! -d Package ]; then
        mkdir Package
    fi
    set -x
    peer lifecycle chaincode package Package/transaction.tar.gz --path ${CC_TRANSACTION_PATH} --lang golang --label transaction_${CC_VERSION} >&log.txt
    cat log.txt
    # peer lifecycle chaincode package Package/market.tar.gz --path ${CC_MARKET_PATH} --lang golang --label market_${CC_VERSION} >&log.txt
    # cat log.txt
    # peer lifecycle chaincode package Package/people.tar.gz --path ${CC_PEOPLE_PATH} --lang golang --label people_${CC_VERSION} >&log.txt
    # cat log.txt
    # peer lifecycle chaincode package Package/revenue.tar.gz --path ${CC_REVENUE_PATH} --lang golang --label revenue_${CC_VERSION} >&log.txt
    # cat log.txt
    set +x
    successln "체인코드 패키징 완료"
}

chaincodeinstall(){
    export CORE_PEER_TLS_ENABLED=true
    echo "Seller에 체인코드 설치"
    setGlobals seller
    peer lifecycle chaincode install Package/transaction.tar.gz >&log.txt
    cat log.txt
    echo "Buyer에 체인코드 설치"
    setGlobals buyer
    peer lifecycle chaincode install Package/transaction.tar.gz >&log.txt
    cat log.txt
    echo "inspector에 체인코드 설치"
    setGlobals inspector
    peer lifecycle chaincode install Package/transaction.tar.gz >&log.txt
    # echo "Seller에 체인코드 설치"
    # setGlobals seller
    # peer lifecycle chaincode install Package/people.tar.gz >&log.txt
    # cat log.txt
    # echo "Buyer에 체인코드 설치"
    # setGlobals buyer
    # peer lifecycle chaincode install Package/people.tar.gz >&log.txt
    # cat log.txt
    # echo "inspector에 체인코드 설치"
    # setGlobals inspector
    # peer lifecycle chaincode install Package/revenue.tar.gz >&log.txt
    # cat log.txt
}

peer version
echo $FABRIC_CFG_PATH

packaging
chaincodeinstall

