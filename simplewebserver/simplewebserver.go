package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Print("Server running.")
	log.Fatal(http.ListenAndServe("localhost:25565", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}

	log.Printf("Request: %s\n", r.Form.Encode())
}
