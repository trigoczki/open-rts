package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	serveApplication()
}

func serveApplication() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHelloOpenRTS).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func handleHelloOpenRTS(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello OpenRTS!"))
}
