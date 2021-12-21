package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Microservices with Go")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Ooops somewthing went wrong", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Helloooo %s", d)
	})

	http.HandleFunc("/exit", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Exit endpoint")
		fmt.Fprintf(rw, "Exit endpoint reached")
	})

	http.ListenAndServe(":8080", nil)
}
