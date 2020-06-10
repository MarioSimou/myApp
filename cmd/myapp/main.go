package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func pong(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "pong\n")
}

func main(){
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	var address = fmt.Sprintf(":%s", port)

	http.HandleFunc("/ping", pong)
	http.HandleFunc("/", home)
	http.HandleFunc("/username", username)
	http.HandleFunc("/new", newRoute)

	log.Printf("The app listens on %s - process id: %s\n", address, uuid.New().String())
	log.Fatalln(http.ListenAndServe(address, nil))
}