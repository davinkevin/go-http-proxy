package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var port string

func serve(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	resp, err := http.Get(q)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)

	w.Write(body)
}

func init() {
	flag.StringVar(&port, "port", "9090", "Port of the proxy web-server")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", serve) // set router
	fmt.Printf("Listening on port : %v", port)
	err := http.ListenAndServe(":"+port, nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
