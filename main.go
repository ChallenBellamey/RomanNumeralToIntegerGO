package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// app

func main() {
	app := mux.NewRouter()
	listen(app)
}

// listen

func listen(router *mux.Router) {
	http.ListenAndServe(":9090", router)
}
