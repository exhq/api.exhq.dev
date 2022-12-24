package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func queryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(res, getshit(req.FormValue("q")))

}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getshit(url string) string {
	gaming := "https://api.song.link/v1-alpha.1/links?url=" + url
	resp, err := http.Get(gaming)
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	return string(html)
}

func main() {
	http.HandleFunc("/youtuber", func(res http.ResponseWriter, req *http.Request) {
		queryParamDisplayHandler(res, req)
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
