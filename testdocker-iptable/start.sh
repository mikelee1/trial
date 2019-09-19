#!/bin/bash

#testalpine2 172.22.0.2
#testalpine3 172.22.0.3 wget 192.168.9.87 9991


#echo `ip a show eth0|grep inet|cut -d/ -f1|awk '{print $2}'`

iptables -t nat -A PREROUTING -d 172.17.0.9 -p tcp --dport 8882 -j DNAT --to-destination 192.168.9.87:9992

iptables -t nat -A POSTROUTING -d 192.168.9.87 -p tcp --dport 9992 -j SNAT --to 172.17.0.9

exec syslogd -n -O -

