package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MarioSimou/cds-library-go/mocks"
	"github.com/google/uuid"
)

func pong(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "pong\n")
}
func welcome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome 40\n")
}

func main(){
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	var msgInterface = mocks.MessageInterface{}
	var address = fmt.Sprintf(":%s", port)
	fmt.Println(msgInterface)

	http.HandleFunc("/ping", pong)
	http.HandleFunc("/", welcome)

	log.Printf("The app listens on %s - process id: %s\n", address, uuid.New().String())
	log.Fatalln(http.ListenAndServe(address, nil))
}