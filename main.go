package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/vinixavier/urlshort/handlers"
)

func main() {
	file := flag.String("file", "", "Specify a YAML file with path to URL.")
	flag.Parse()

	yamlHandler, err := handlers.New().YAML(file, nil)
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr:    ":8000",
		Handler: &yamlHandler,
	}
	fmt.Printf("Starting server on %s", server.Addr)

	server.ListenAndServe()

}
