from youtuber import youtuber
from flask import Flask
app = Flask(__name__)


@app.route('/youtuber/<link>')
def show_user_profile(link):
    x = youtuber(link)
    return x


if __name__ == '__main__':
    from waitress import serve
    serve(app, host="0.0.0.0", port=5000)
