#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-buyer.0
#

# This is a collection of bash functions used by different scripts

# imports
. scripts/utils.sh

export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/tlsca/tlsca.example.com-cert.pem
export PEER0_seller_CA=${PWD}/organizations/peerOrganizations/seller.example.com/tlsca/tlsca.seller.example.com-cert.pem
export PEER0_buyer_CA=${PWD}/organizations/peerOrganizations/buyer.example.com/tlsca/tlsca.buyer.example.com-cert.pem
export PEER0_inspector_CA=${PWD}/organizations/peerOrganizations/inspector.example.com/tlsca/tlsca.inspector.example.com-cert.pem
export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.crt
export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.key

# Set environment variables for the peer org
setGlobals() {
  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  infoln "Using organization ${USING_ORG}"
  if [ $USING_ORG == "seller" ]; then
    export CORE_PEER_LOCALMSPID="sellerMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_seller_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/seller.example.com/users/Admin@seller.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7051
  elif [ $USING_ORG == "buyer" ]; then
    export CORE_PEER_LOCALMSPID="buyerMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_buyer_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/buyer.example.com/users/Admin@buyer.example.com/msp
    export CORE_PEER_ADDRESS=localhost:9051

  elif [ $USING_ORG == "inspector" ]; then
    export CORE_PEER_LOCALMSPID="inspectorMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_inspector_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/inspector.example.com/users/Admin@inspector.example.com/msp
    export CORE_PEER_ADDRESS=localhost:11051
  else
    infoln "error"
    errorln "ORG Unknown"
  fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

# Set environment variables for use in the CLI container
setGlobalsCLI() {
  setGlobals $1

  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  if [ $USING_ORG == "seller" ]; then
    export CORE_PEER_ADDRESS=peer0.seller.example.com:7051
  elif [ $USING_ORG == "buyer" ]; then
    export CORE_PEER_ADDRESS=peer0.buyer.example.com:9051
  elif [ $USING_ORG == "inspector" ]; then
    export CORE_PEER_ADDRESS=peer0.inspector.example.com:11051
  else
    errorln "ORG Unknown"
  fi
}

# parsePeerConnectionParameters $@
# Helper function that sets the peer connection parameters for a chaincodes
# operation
parsePeerConnectionParameters() {
  PEER_CONN_PARMS=()
  PEERS=""
  while [ "$#" -gt 0 ]; do
    setGlobals $1
    PEER="peer0.$1"
    ## Set peer addresses
    if [ -z "$PEERS" ]
    then
	PEERS="$PEER"
    else
	PEERS="$PEERS $PEER"
    fi
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" --peerAddresses $CORE_PEER_ADDRESS)
    ## Set path to TLS certificate
    CA=PEER0_ORG$1_CA
    TLSINFO=(--tlsRootCertFiles "${!CA}")
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" "${TLSINFO[@]}")
    # shift by one to get to the next organization
    shift
  done
}

verifyResult() {
  if [ $1 -ne 0 ]; then
    fatalln "$1"
  fi
}
