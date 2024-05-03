import os

from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/', methods=['GET'])
def health_check():
    return jsonify({
        'message': 'Server is up',
        'status': 200,
    })

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=1234, debug=True)