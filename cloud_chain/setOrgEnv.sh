#!/bin/bash

if [ "$1" = "seller" ]; then
  role="seller"
  port=7051
elif [ "$1" = "buyer" ]; then
  role="buyer"
  port=9051
elif [ "$1" = "inspector" ]; then
  role="inspector"
  port=11051
else
  echo "Invalid role specified. Please provide either 'seller', 'buyer', or 'inspector' as an argument."
fi

export FABRIC_CFG_PATH="${PWD}/configtx"
export VERBOSE=false
export PATH="${PWD}/../bin:$PATH"
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="${role}MSP"
export CORE_PEER_TLS_ROOTCERT_FILE="${PWD}/organizations/peerOrganizations/${role}.example.com/peers/peer0.${role}.example.com/tls/ca.crt"
export CORE_PEER_MSPCONFIGPATH="${PWD}/organizations/peerOrganizations/${role}.example.com/users/Admin@${role}.example.com/msp"
export CORE_PEER_ADDRESS=localhost:${port}
export ORDERER_CA="${PWD}/organizations/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem"
export ORDERER_ADDRESS=localhost:7050
