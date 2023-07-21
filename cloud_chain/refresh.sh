#!/bin/bash

./AAS.sh down
./AAS.sh crypto
./AAS.sh up -s couchdb
./AAS.sh set
./AAS.sh install