#!/bin/bash

#testalpine2 172.22.0.2
#testalpine3 172.22.0.3 wget 192.168.9.87 9991

#获取ipv4的地址
IPADDR=`ip a show eth0|grep inet|cut -d/ -f1|head -1|awk '{print $2}'`
#从8882转出到9992
iptables -t nat -A PREROUTING -d $IPADDR -p tcp --dport 8882 -j DNAT --to-destination 192.168.9.87:30021
#从9992转回到8882
iptables -t nat -A POSTROUTING -d 192.168.9.87 -p tcp --dport 30021 -j SNAT --to $IPADDR

exec syslogd -n -O -




