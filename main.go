package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Scheme = "https"
	r.URL.Host = r.Host
	http.Redirect(w, r, r.URL.String(), http.StatusMovedPermanently)
	return
}
