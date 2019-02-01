package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("usage: request to http://localhost:8080")

	http.HandleFunc("/", handlerFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s received from: %s\n", r.Method, r.URL)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("This is the response for your %s request.\n", r.Method)))

	return
}
