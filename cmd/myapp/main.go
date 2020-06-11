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
func welcome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome\n")
}
func showEnv(w http.ResponseWriter, r *http.Request){
	var envName = r.URL.Query().Get("q")
	if v := os.Getenv(envName); envName != "" && v != ""  {
		fmt.Fprintf(w, fmt.Sprintf("name: %s", v))
	} else {
		fmt.Fprintf(w, "No match for this variable")		
	}
}

func main(){
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	var address = fmt.Sprintf(":%s", port)

	http.HandleFunc("/ping", pong)
	http.HandleFunc("/", welcome)
	http.HandleFunc("/env", showEnv)

	log.Printf("The app listens on %s - process id: %s\n", address, uuid.New().String())
	log.Fatalln(http.ListenAndServe(address, nil))
}