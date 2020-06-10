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

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome home\n")
}

func username(w http.ResponseWriter, r *http.Request){
	var username = r.URL.Query().Get("name")
	if username == "" {
		fmt.Fprintf(w, "I don't know your name")
	} else {
		fmt.Fprintf(w, "Hello %s\n", username)		
	}
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

	log.Printf("The app listens on %s - process id: %s\n", address, uuid.New().String())
	log.Fatalln(http.ListenAndServe(address, nil))
}