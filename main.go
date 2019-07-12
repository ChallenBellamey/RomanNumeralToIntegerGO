package main

import (
	"encoding/json"
	"fmt"
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

	router.HandleFunc("/api/roman-numerals-to-integer", rntiRouter).Methods("GET")

	router.HandleFunc("/api", apiRouter).Methods("GET")

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handle400(w, error{Code: 400, Message: "Invalid url."})
	}).Methods("GET")

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handle405(w, error{Code: 405, Message: "Method not allowed."})
	}).Methods("POST", "PATCH", "DELETE")
}

func apiRouter(writer http.ResponseWriter, request *http.Request) {
	getInfo(writer, request)
}

func rntiRouter(writer http.ResponseWriter, request *http.Request) {
	if rn, boolean := request.URL.Query()["rn"]; boolean {
		fmt.Println(rn)
	} else {
		getInfo(writer, request)
	}
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

// errors

type error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func handle400(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(err)
}

func handle405(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusMethodNotAllowed)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(err)
}
