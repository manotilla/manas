#!/bin/bash

CONTAINERS=$(docker ps -q)
for container in $CONTAINERS; do

	CNT_PID=$(docker inspect $container |grep -i '"Pid":'|tr -d " "|awk '{split($0,a,":"); print a[2]}'|tr -d ',')
	IPC=$(ls -la "/proc/$CNT_PID/ns" 2>err.log|awk '{split($0,a,"->"); print a[2]}'|grep "ipc")

	echo $CNT_PID
	echo $IPC
	redis-cli -n 0 HMSET "container-$container" "ipc" $IPC

done