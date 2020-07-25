
#!/bin/bash

ALL_PIDS=$(ps -ef|awk '{print $2}')

for pid in $ALL_PIDS; do
	CMD_LINE=$(cat "/proc/$pid/cmdline" 2>/dev/null)
	IPC=$(ls -la '/proc/$pid/ns' 2>/dev/null|awk '{split($0,a,"->"); print a[2]}'|grep "ipc")
	redis-cli -n 0 HMSET $IPC $pid $CMD_LINE

done


