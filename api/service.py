from flask import Flask, jsonify, request, render_template
from proc_info import *
import json

app = Flask(__name__)

@app.route("/dashboard", methods=['GET'])
def dashboard():
    response = generateProcResponse()
    return render_template("dashboard.html", processes=response)

if __name__ == "__main__":
    app.run("0.0.0.0", port=5000, debug=True)
