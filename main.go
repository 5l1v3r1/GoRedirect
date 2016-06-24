package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Scheme = "https://"
	http.Redirect(w, r, "", http.StatusMovedPermanently)
	return
}
