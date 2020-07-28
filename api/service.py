from flask import Flask, jsonify, request
from proc_info import *
import json

app = Flask(__name__)
@app.route("/processes", methods=['GET'])
def processes():

    container_pid = request.args.getlist('container_pid', type=str)
    data = getIPCProcesses(ipc)

    response_map = {"processes": data}
    return jsonify(response_map)


if __name__ == "__main__":
    app.run("0.0.0.0", port=5000, debug=True)
