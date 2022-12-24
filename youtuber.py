import requests
import validators


def youtuber(x):
    x = x.replace("youtuber/", "")
    if validators.url(x):

        lmao = requests.get(url="https://song.link/"+x).text
        actuallink = str(lmao[lmao.find("www.youtube.com/watch?")
                         :lmao.find("www.youtube.com/watch?")+35])
        title = str(lmao[lmao.find("<title>"):lmao.find("</titl")])

        head = f"""<title>MY API IS GOOD</title>
    <meta content="youtuber :fire:" property="og:title" />
    <meta content="{title}" property="og:description" />
    <meta content="{actuallink}" property="og:url" />
"""
        return head + title.replace("title", "").replace("<>", "") + "\n" + actuallink
    else:
        return "NOT A URL BRO"
