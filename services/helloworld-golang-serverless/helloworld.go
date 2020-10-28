package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloworld!\n")
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PT_CONTAINER_PORT")
	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
