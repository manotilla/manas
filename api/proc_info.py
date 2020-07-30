import redis
import os

REDIS_HOST=os.environ["REDIS_HOST"]

redis_conn = redis.Redis(host=REDIS_HOST, port=6379, db=0, charset="utf-8", decode_responses=True)

def getContainerIPC(container_id):
    search_str="container-"+container_id

    data = redis_conn.hget(search_str, "ipc")

    return data

def getIPCProcesses(ipc):

    data = redis_conn.lrange(ipc, 0, 5)

    return data

def generateProcResponse():

    return ["29e4110398ac", "51c396f0d899"]



def generateProcResponse():

    containers = ["29e4110398ac", "51c396f0d899"]
    response = []

    for container in containers:
        ipc = getContainerIPC(container)
        tmpJson = {"ipc": ipc, "processes":getIPCProcesses(ipc)}
        response.append(tmpJson)

    return response

