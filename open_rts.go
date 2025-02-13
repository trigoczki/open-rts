package main

import (
	"encoding/json"
	"net/http"
	"open-rts/cache"

	"github.com/gorilla/mux"
)

func main() {
	cache.Initialize()
	serveApplication()
}

func serveApplication() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHelloOpenRTS).Methods(http.MethodGet)
	router.HandleFunc("/resources", handleGetResources).Methods(http.MethodGet)
	router.HandleFunc("/buildings", handleGetBuildings).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func handleGetBuildings(w http.ResponseWriter, r *http.Request) {
	buildings := cache.GetInstance().GetBuildings()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buildings)
}

func handleHelloOpenRTS(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello OpenRTS!"))
}

func handleGetResources(w http.ResponseWriter, r *http.Request) {
	resources := cache.GetInstance().GetResources()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}
