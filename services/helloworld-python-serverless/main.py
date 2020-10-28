import os

from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello_world():

    with open('/settings/settings.json', 'r') as fh:
        print(fh.read())

    return 'Hello World - from python'


if __name__ == "__main__":
    app.run(debug=True, host='0.0.0.0', port=int(os.environ['PT_CONTAINER_PORT']))
