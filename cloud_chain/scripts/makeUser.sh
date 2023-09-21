#!/bin/bash
PROJECT_PATH=/Users/jeho/Desktop/selab/graduate_project/blockchain-repo/cloud_chain
ORG=$1
ID=$2
PW=$3
echo $ORG $ID $PW

export FABRIC_CA_CLIENT_HOME=${PROJECT_PATH}/organizations/peerOrganizations/${ORG}.example.com/

fabric-ca-client register --caname ca-${ORG} --id.name ${ID} --id.secret ${PW} --id.type client --tls.certfiles "${PROJECT_PATH}/organizations/fabric-ca/${ORG}/ca-cert.pem"
# sleep 1
fabric-ca-client enroll -u http://${ID}:${PW}@localhost:8054 --caname ca-${ORG} -M "${PROJECT_PATH}/organizations/peerOrganizations/${ORG}.example.com/users/${ID}@${ORG}.example.com/msp" --tls.certfiles "${PROJECT_PATH}/organizations/fabric-ca/${ORG}/ca-cert.pem"