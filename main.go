package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	var p string

	flag.StringVar(&p, "p", "8000", "application port")

	// parse flags
	flag.Parse()

	h := http.NewServeMux()

	h.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("pong"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("server internal error"))
		}
	})

	log.Printf("[INFO] starting application on port %s", p)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", p), h); err != nil {
		log.Fatalln(err)
	}
}
