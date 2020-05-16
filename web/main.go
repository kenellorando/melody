package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/kenellorando/clog"
)

func main() {
	// Handle routes
	r := mux.NewRouter()

	//
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css/"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js/"))))
	r.HandleFunc("/", ServeRoot).Methods("GET")

	// Start server
	clog.Info("main", fmt.Sprintf("Starting webserver on port <%s>.", ":8900"))
	clog.Fatal("main", "Server failed to start!", http.ListenAndServe(":8900", r))
}

// ServeRoot - serve the homepage
func ServeRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, path.Dir("./public/index.html"))
}
