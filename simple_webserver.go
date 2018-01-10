package main

import (
	"fmt"
	//"html"
	"log"
	"net/http"
	"io/ioutil"
)

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("./html"))))

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query()["time"]
		locationJson := GetRequest("http://api.flutrack.org/?time=" + string(param[0]))
		fmt.Fprintf(w, locationJson)
	})

	log.Fatal(http.ListenAndServe(":8010", nil))
}

func GetRequest(address string) string {
	resp, err := http.Get(address)
	if err != nil {
		log.Println("Request error")
		return "Request error"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}