from bs4 import BeautifulSoup
import flask
app = flask.Flask(__name__)


@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def catch_all(path):
    if bool(BeautifulSoup(path, "html.parser").find()):
        return "kill yourself"
    else:
        return path


if __name__ == '__main__':
    app.run()
