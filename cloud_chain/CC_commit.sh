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
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt

    setGlobals buyer
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt

    setGlobals inspector
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt
    
    # Peer 0에 대한 체인코드 배포
    infoln "채널 vehicles 커밋 - Peer 0"
    peer lifecycle chaincode commit -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/seller.pnu.cse/peers/peer0.seller.pnu.cse/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/buyer.pnu.cse/peers/peer0.buyer.pnu.cse/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/inspector.pnu.cse/peers/peer0.inspector.pnu.cse/tls/ca.crt"


    infoln "채널 vehicles 확인"
    peer lifecycle chaincode querycommitted --channelID ${CHANNEL_NAME} --name ${CC_NAME} --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"

    CC_NAME="inspection"
    setGlobals seller
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)

    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt

    setGlobals buyer
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt

    setGlobals inspector
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt

    # Peer 0에 대한 체인코드 배포
    infoln "채널 vehicles 커밋 - Peer 0"
    peer lifecycle chaincode commit -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --sequence 1 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/seller.pnu.cse/peers/peer0.seller.pnu.cse/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/buyer.pnu.cse/peers/peer0.buyer.pnu.cse/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/inspector.pnu.cse/peers/peer0.inspector.pnu.cse/tls/ca.crt" 

    infoln "채널 vehicles 확인"
    peer lifecycle chaincode querycommitted --channelID ${CHANNEL_NAME} --name ${CC_NAME} --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"

}
getcheck2(){
    export CC_END_POLICY="--signature-policy OR('sellerMSP.peer','buyerMSP.peer','inspectorMSP.peer')" 
    CC_NAME="transaction"
    CC_VERSION="1.0"
    setGlobals2 seller
    echo $CHANNEL_NAME
    echo $CC_END_POLICY
    echo ${CC_NAME}
    
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    echo ${CC_PACKAGE_ID}
    echo $CC_PACKAGE_ID
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt

    setGlobals2 buyer
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt

    setGlobals2 inspector
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt
    
     # Peer 1에 대한 체인코드 배포
    infoln "채널 vehicles 커밋 - Peer 1"
    peer lifecycle chaincode commit -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --peerAddresses localhost:7151 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/seller.pnu.cse/peers/peer1.seller.pnu.cse/tls/ca.crt" --peerAddresses localhost:9151 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/buyer.pnu.cse/peers/peer1.buyer.pnu.cse/tls/ca.crt" --peerAddresses localhost:11151 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/inspector.pnu.cse/peers/peer1.inspector.pnu.cse/tls/ca.crt"

    infoln "채널 vehicles 확인"
    peer lifecycle chaincode querycommitted --channelID ${CHANNEL_NAME} --name ${CC_NAME} --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"

    CC_NAME="inspection"
    setGlobals2 seller
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)

    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt

    setGlobals2 buyer
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt

    setGlobals2 inspector
    peer lifecycle chaincode queryinstalled >&log.txt
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json
    CC_PACKAGE_ID=$(sed -n "/${CC_NAME}_${CC_VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" log.txt)
    peer lifecycle chaincode approveformyorg -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --package-id $CC_PACKAGE_ID --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"
    peer lifecycle chaincode checkcommitreadiness --channelID ${CHANNEL_NAME} ${CC_END_POLICY} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --output json

    cat log.txt

    # Peer 1에 대한 체인코드 배포
    infoln "채널 vehicles 커밋 - Peer 1"
    peer lifecycle chaincode commit -o localhost:7050 ${CC_END_POLICY} --ordererTLSHostnameOverride orderer.pnu.cse --channelID ${CHANNEL_NAME} --name ${CC_NAME} --version 1.0 --sequence 2 --tls --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem" --peerAddresses localhost:7151 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/seller.pnu.cse/peers/peer1.seller.pnu.cse/tls/ca.crt" --peerAddresses localhost:9151 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/buyer.pnu.cse/peers/peer1.buyer.pnu.cse/tls/ca.crt" --peerAddresses localhost:11151 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/inspector.pnu.cse/peers/peer1.inspector.pnu.cse/tls/ca.crt"
    
    infoln "채널 vehicles 확인"
    peer lifecycle chaincode querycommitted --channelID ${CHANNEL_NAME} --name ${CC_NAME} --cafile "${PWD}/organizations/ordererOrganizations/pnu.cse/orderers/orderer.pnu.cse/msp/tlscacerts/tlsca.pnu.cse-cert.pem"


}

   
    
getcheck
getcheck2


