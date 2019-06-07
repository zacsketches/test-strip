package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./tmp")))

	fmt.Println("listening on port 9000")
	err := http.ListenAndServe("localhost:9000", r)
	if err != nil {
		log.Fatal(err)
	}
}
