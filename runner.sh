#!/bin/bash

# Fail on error
set -e

pem_file_path="~/Downloads/graphDbIreland.pem"
port="20301"
rep=100

ip=$1
if [ -z "$1" ]
then
    echo "Provide the Ip address"
    exit
fi

rm -f algo
GOARCH=arm64 make algo
scp -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ./algo ubuntu@$ip:~/
scp -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ./*.csv ubuntu@$ip:~/
scp -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ./queries.txt ubuntu@$ip:~/
scp -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ./*.sh ubuntu@$ip:~/
echo "Running graph algorithm service"
ssh -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ubuntu@$ip "nohup ./run_variations_s3.sh </dev/null &"

scp -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ubuntu@$ip:~/s3*.txt .
