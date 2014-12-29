package main

import (
	"fmt"
	nw "github.com/lonnc/golang-nw"
	"net/http"
	"log"
)

func nwExample() {

	// Setup our handler
	http.HandleFunc("/", hello2)

	// Create a link back to node-webkit using the environment variable
	// populated by golang-nw's node-webkit code
	nodeWebkit, err := nw.New()
	if err != nil {
		panic(err)
	}


	// Pick a random localhost port, start listening for http requests using default handler
	// and send a message back to node-webkit to redirect
	if err := nodeWebkit.ListenAndServe(nil); err != nil {
		log.Fatalln(err)
	}
}

func hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from golang.")
}
