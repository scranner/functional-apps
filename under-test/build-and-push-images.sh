#!/bin/bash

cd api-gateway/subprojects || exit
./build.sh "$1"
./push-images.sh "$1"
cd ../.. || exit

cd database-lookup/subprojects || exit
./build.sh "$1"
./push-images.sh "$1"
