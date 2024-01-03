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
echo $ip

rm graph_algorithm_service
GOARCH=arm64 make graph_algorithm_service
scp -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ./graph_algorithm_service ubuntu@$ip:~/
ssh -o StrictHostKeyChecking=accept-new -i $pem_file_path\
    ubuntu@$ip "./graph_algorithm_service --address localhost:$port --repetitions $rep | cat >> client.log"
