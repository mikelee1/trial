#!/usr/bin/env bash

##query操作
#starttime=`date +%s`
#echo $starttime
##100=4s
##1000=40s
#for ((i=0;i<100;i++))
#do
#    curl -H "Content-Type:application/json" -X POST --data '{"channel":"mychannel","ccName":"usercc","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJKdGkiOiI2In0.8cFCiqzh4eUAHINhGHi9QAkahWMekQcA1Qs5OV13Ifk","fcn":"query","args":["a"]}' http://192.168.9.21:8089/chaincode/query > /dev/null 2>&1
#done
#endtime=`date +%s`
#echo $endtime
#
#echo "本次运行时间： "$((endtime-starttime))"s"



#invoke操作
starttime=`date +%s`
echo $starttime
#10=22s
#20=44s
for ((i=0;i<20;i++))
do
    curl -H "Content-Type:application/json" -X POST --data '{"channel":"mychannel","ccName":"usercc","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJKdGkiOiI2In0.8cFCiqzh4eUAHINhGHi9QAkahWMekQcA1Qs5OV13Ifk","fcn":"invoke","args":["a","b","2"]}' http://192.168.9.21:8089/chaincode/invoke > /dev/null 2>&1
done
endtime=`date +%s`
echo $endtime

echo "本次运行时间： "$((endtime-starttime))"s"


