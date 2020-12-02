package main

import (
	"Go-000/Week02/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	result := service.GetUser()
	fmt.Fprintf(w, result)
}

func main() {
	http.HandleFunc("/", getUser)
	http.ListenAndServe(":8000", nil)
}
