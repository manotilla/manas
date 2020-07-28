import redis
import os

REDIS_HOST=os.environ["REDIS_HOST"]

redis_conn = redis.Redis(host=REDIS_HOST, port=6379, db=0, charset="utf-8", decode_responses=True)

def getContainerIPC(container_pid):
    search_str="container-"+container_pid

    data = redis_conn.hget(search_str, "ipc")

    return data

def getIPCProcesses(ipc):

    data = redis_conn.lrange(ipc, 0, 20)

    return data

