from flask import Flask
import balls
app = Flask(__name__)


@app.route('/balls')
def index():
    return balls.dickandballs()


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
