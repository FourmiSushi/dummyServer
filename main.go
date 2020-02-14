package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("usage: request to http://localhost:8080")

	http.HandleFunc("/", handlerFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----")

	log.Printf("%s received from: %s\n", r.Method, r.RemoteAddr)
	log.Printf("Content-Type: %s", r.Header.Get("content-type"))
	log.Printf("URI: %s", r.RequestURI)

	switch r.Method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("couldnt read body from request. is it nil?")
			break
		}
		log.Printf("%s\n", body)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("This is the response for your %s request.\n", r.Method)))

	return
}
