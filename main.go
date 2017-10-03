package main

import (
	"fmt"

	"github.com/GanymedeNil/shorturl/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func GetSurl(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	ee := queryParams["u"][0]
	fmt.Fprint(w, data.SetURL(ee))

}

func getOurl(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	ee := queryParams["u"][0]
	fmt.Fprint(w, data.GetURL(ee))

}

func main() {
	r := mux.NewRouter()
	// Attach an elegant path with handler
	r.HandleFunc("/api/v1/getsurl", GetSurl)
	r.HandleFunc("/api/v1/getourl", getOurl)
	r.Queries("u")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
