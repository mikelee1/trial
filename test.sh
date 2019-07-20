#!/usr/bin/env bash

#for i in `docker images|grep hyperledger|grep amd64-l|awk '{print $1 aaa $3}'`;
#do
#  echo $i
#done

for i in  `ls hyperledger`
do
    echo $i
    docker load -i $i
done

