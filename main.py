from bs4 import BeautifulSoup
import flask
from youtuber import youtuber
app = flask.Flask(__name__)


@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def catch_all(path):

    if bool(BeautifulSoup(path, "html.parser").find()):
        return "kill yourself"
    else:
        strpath = str(path).split("/")
        if len(path) == 0:
            return "ill add a api list at some point lol"
        if path[0] == "youtuber":
            x = youtuber(str(path).replace("youtuber/", ""))
            return x
        else:
            return "unknown api: " + path[0]


if __name__ == '__main__':
    from waitress import serve
    serve(app, host="0.0.0.0", port=5000)
