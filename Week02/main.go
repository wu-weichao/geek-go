package main

import (
	"log"
	"net/http"
	_ "week02/dao"
	"week02/service"
)

func main() {
	http.HandleFunc("/search", service.GetUser)
	log.Print("starting")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
