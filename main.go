package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func queryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	print(req.FormValue("q"))

	io.WriteString(res, getjson(req.FormValue("q")))
}

func getjson(link string) string {
	resp, err := http.Get("https://api.song.link/v1-alpha.1/links?url=" + link)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
	http.HandleFunc("/youtuber", func(res http.ResponseWriter, req *http.Request) {
		queryParamDisplayHandler(res, req)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
