#!/bin/bash

# Fail on error
set -e

pem_file_path="~/Downloads/graphDbIreland.pem"
port="20301"
rep=1000

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
echo "Running graph algorithm service"
ssh -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ubuntu@$ip "./algo --address localhost:$port --repetitions $rep 2> client.log"
ssh -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ubuntu@$ip "pkill access"

dirName="test_$(date +%s)"
mkdir $dirName
scp -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ubuntu@$ip:~/server.log ./$dirName/
scp -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ubuntu@$ip:~/client.log ./$dirName/
