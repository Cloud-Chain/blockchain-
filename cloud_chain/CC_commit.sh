#!/bin/bash

source scripts/utils.sh
. scripts/envVar.sh
. scripts/ccutils.sh
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CC_VERSION="1.0"
export CHANNEL_NAME="vehicles"
getcheck(){
    export CC_END_POLICY="--signature-policy OR('sellerMSP.peer','buyerMSP.peer','inspectorMSP.peer')" 
    CC_NAME="transaction"
    CC_VERSION="1.0"
    setGlobals seller
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.example.com --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --output json

    cat log.txt

    setGlobals buyer
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.example.com --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --output json

    cat log.txt

    setGlobals inspector
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.example.com --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --output json

    cat log.txt
    
    infoln "채널 vehicles 커밋"
    peer lifecycle chaincode commit -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.example.com --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/seller.example.com/peers/peer0.seller.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/buyer.example.com/peers/peer0.buyer.example.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/inspector.example.com/peers/peer0.inspector.example.com/tls/ca.crt"
    infoln "채널 vehicles 확인"
    peer lifecycle chaincode querycommitted --channelID ${CHANNEL_NAME} --name ${CC_NAME} --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
     
}


getcheck



