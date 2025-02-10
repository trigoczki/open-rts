package main

import (
	"fmt"
	"net/http"
	ruleset "open-rts/ruleset/resource"

	"github.com/gorilla/mux"
)

func main() {
	resources, err := ruleset.ParseResources("ruleset/resource/resources_ruleset.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(resources)

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
