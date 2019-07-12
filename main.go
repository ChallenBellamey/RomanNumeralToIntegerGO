package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// app

func main() {
	app := mux.NewRouter()
	routes(app)
	listen(app)
}

// listen

func listen(router *mux.Router) {
	http.ListenAndServe(":9090", router)
}

// routes

func routes(router *mux.Router) {
	router.HandleFunc("/api", getInfo).Methods("GET")
}

// api-controller

func getInfo(writer http.ResponseWriter, request *http.Request) {
	message := "Hello! Please attach a query to your get request to receive a Roman numerals conversion, i.e. " + request.Host + "/api/roman-numerals-to-integer?rn=CC"

	type Response struct {
		Message string `json:"message"`
	}

	response := Response{Message: message}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
