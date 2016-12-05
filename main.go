package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)

	// Must set timeouts otherwise we'll inevitably run out of file descriptors
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe(":80", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Scheme = "https"
	r.URL.Host = r.Host
	http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
	return
}
