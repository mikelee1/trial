#!/usr/bin/env bash


#invoke操作
starttime=`date +%s`
echo $starttime
#10=22s
#20=44s
for ((i=0;i<2000;i++))
do
    ./client
done
endtime=`date +%s`
echo $endtime

echo "本次运行时间： "$((endtime-starttime))"s"