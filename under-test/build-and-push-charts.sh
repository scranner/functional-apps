#!/bin/bash

cd api-gateway/subprojects || exit
./push-charts.sh
cd ../api-gateway || exit
helm dep update
helm push . -f dev
cd ../.. || exit

cd database-lookup/subprojects || exit
./push-charts.sh
cd ../database-lookup || exit
helm dep update
helm push . -f dev