package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}
	path := "welcome"
	if isPreferredCy(r.Header.Get("Accept-Language")) {
		path = "croeso"
	}
	http.Redirect(w, r, fmt.Sprintf("http://localhost:%d/%s", port, path), http.StatusFound)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}
	acceptLanguageHeader := r.Header.Get("Accept-Language")
	preferCy := isPreferredCy(acceptLanguageHeader)
	fmt.Fprintf(w, "Accept-Language: %s\nPrefer cy: %t\n", acceptLanguageHeader, preferCy)
}

func croesoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Croeso i'r tudalen Cymraeg\n")
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Welcome to the English page\n")
}
