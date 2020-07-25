#!/bin/bash

while true
do
	CONTAINER_PID=$(docker ps -q |xargs docker inspect |grep -i '"Pid":'|tr -d " "|awk '{split($0,a,":"); print a[2]}'|tr -d ',')
	IPC=$(ls -la "/proc/$CONTAINER_PID/ns" 2>err.log|awk '{split($0,a,"->"); print a[2]}'|grep "ipc")
	echo "$IPC"
	echo "$CONTAINER_PID"

	redis-cli -n 0 HMSET "container-$CONTAINER_PID" "ipc" $IPC
done