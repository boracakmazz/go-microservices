package main

import (
	"log"
	"main/handlers"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stderr, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gg := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gg)

	http.ListenAndServe(":8080", sm)
}
