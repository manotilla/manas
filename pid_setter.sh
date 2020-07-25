#!/bin/bash

ALL_PIDS=$(ps -ef|awk '{print $2}')

while true
do
	for pid in $ALL_PIDS; do
		echo "PID: $pid"
		CMD_LINE=$(cat "/proc/$pid/cmdline"|tr -d '\0')

		IPC=$(ls -la "/proc/$pid/ns" 2>err.log|awk '{split($0,a,"->"); print a[2]}'|grep "ipc")

		if [[ "$IPC" == *"ipc:["* ]]
		then
			echo "$IPC"
			echo "$CMD_LINE"
			redis-cli -n 0 HMSET $pid $IPC "process-$CMD_LINE"
		else
			echo "empty IPC output"
		fi
	done

done