package main

import (
    "fmt"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the crapiest URLShortener!")
}

func RedirectLink(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Location", "http://google.es")
    w.WriteHeader(301)
}

func LinkStats(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to LinkStats!")
}

func NewLink(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to NewLink!")
}