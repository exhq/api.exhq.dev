package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func queryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	res.Header().Set("Access-Control-Allow-Origin", "*")
	print(req.FormValue("q"))

	io.WriteString(res, "gay")

}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type Foo struct {
	Bar string
}

func main() {
	//http.HandleFunc("/youtuber", func(res http.ResponseWriter, req *http.Request) {
	//	queryParamDisplayHandler(res, req)
	//})
	//log.Fatal(http.ListenAndServe(":5000", nil))
	url := "https://api.song.link/v1-alpha.1/links?url=https://open.spotify.com/track/4OHAGy4roqG6Aog8XrTOez?si=P_OEMTCYR6-84UA8LKxKUQ"
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
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
	fmt.Printf("%s\n", html)
}
