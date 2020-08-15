import redis
import os

class Proc(object):

    def __init__(self):
        REDIS_HOST=os.environ["REDIS_HOST"]
        self.redis_conn = redis.Redis(host=REDIS_HOST, port=6379, db=0, charset="utf-8", decode_responses=True)

    def getContainerIPC(self, container_id):
        search_str="container-"+container_id

        data = self.redis_conn.hget(search_str, "ipc")

        return data

    def getIPCProcesses(self, ipc):

        data = self.redis_conn.lrange(ipc, 0, 10)

        return data


    def generateProcResponse(self, ):

        containers = ["c2efb9b3a283", "ec6a3a1f7e1b"]
        response = []

        for container in containers:
            ipc = self.getContainerIPC(container)
            tmpJson = {"container":container, "ipc": ipc, "processes":self.getIPCProcesses(ipc)}
            response.append(tmpJson)

        return response

