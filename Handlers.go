package main

import (
    "fmt"
    "net/http"
    "strings"

)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the crapiest URLShortener!")
}

func RedirectLink(w http.ResponseWriter, r *http.Request) {
    longURL := StorageLookup(strings.Trim(r.URL.Path,"/"))
    if longURL == ""{
        http.Error(w, "URL, not found, buddie :S", http.StatusNotFound)
        //w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        //w.WriteHeader()
    } else {
        w.Header().Set("Location", longURL)
        w.WriteHeader(301)    
    }
}

func LinkStats(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to LinkStats2!")
}

func NewLink(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to NewLink!")
}
func RedirectLink2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome %s!",r.URL.Path)
    // w.Header().Set("Location", "http://google.es")
    // w.WriteHeader(301)
}