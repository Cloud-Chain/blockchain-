#!/bin/bash
PROJECT_PATH=/home/jeho/blockchain-repo/cloud_chain
ORG=$1
ID=$2
PW=$3
echo $ORG $ID $PW
seller="seller"
buyer="buyer"
inspector="inspector"

export FABRIC_CA_CLIENT_HOME=${PROJECT_PATH}/organizations/peerOrganizations/${ORG}.pnu.cse/
export PATH=$PATH:/home/jeho/blockchain-repo/bin
echo $FABRIC_CA_CLIENT_HOME
fabric-ca-client register --caname ca-${ORG} --id.name ${ID} --id.secret ${PW} --id.type client --tls.certfiles "${PROJECT_PATH}/organizations/fabric-ca/${ORG}/ca-cert.pem"
# sleep 1
if [ $ORG == $seller ]
then
	port=7054
elif [ $ORG == $buyer ]
then
	port=8054
else
	port=9054
fi
echo $ORG $seller $buyer $port

fabric-ca-client enroll -u http://${ID}:${PW}@localhost:${port} --caname ca-${ORG} -M "${PROJECT_PATH}/organizations/peerOrganizations/${ORG}.pnu.cse/users/${ID}@${ORG}.pnu.cse/msp" --tls.certfiles "${PROJECT_PATH}/organizations/fabric-ca/${ORG}/ca-cert.pem"
