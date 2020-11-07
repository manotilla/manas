import docker
import logging
import os

try:
    if os.environ["LOG_LEVEL"] == "DEBUG":
        log_level = logging.DEBUG
    else:
        log_level = logging.INFO
except:
    log_level = logging.INFO

logging.basicConfig(level=log_level)

class Containers(object):
    def __init__(self):
        self.container_client = docker.from_env()

    def detect_host_ip(self, checked_ip):
        containers = self.container_client.containers.list()
        for container in containers:

            try: 
                source_ip = container.attrs["NetworkSettings"]["Networks"]["bridge"]["IPAddress"]
                image_id = container.attrs['Config']['Image']
                container_id = container.attrs['Id'][0:4]
                if source_ip == checked_ip:
                    logging.info({"container_id": container_id, "image_id": image_id, "source_ip": checked_ip})
                    return {"container_id": container_id, "image_id": image_id , "source_ip": checked_ip}
                else:
                    continue
            except Exception as exp:
                logging.error(exp)
