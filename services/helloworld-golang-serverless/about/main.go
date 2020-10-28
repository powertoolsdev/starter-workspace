package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page\n")
}

func main() {
	log.Print("helloworld: starting about server...")

	http.HandleFunc("/", handler)

	port := os.Getenv("PT_CONTAINER_PORT")
	log.Printf("helloworld-about: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
