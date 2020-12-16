from flask import Flask, abort, request
from factorial import factorial

app = Flask(__name__)


@app.route("/")
def root():
    result = factorial(request.args.get('x'))
    if result is None:
        abort(400, "Invalid Parameter")
    return result


@app.route("/live")
def live():
    return '{"live":"OK"}'


@app.route("/ready")
def ready():
    return '{"ready":"OK"}'


if __name__ == "__main__":
    app.run(host='0.0.0.0', port=80)
