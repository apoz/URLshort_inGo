package main

import (
    "fmt"
    "net/http"
    "strings"
    "encoding/json"
    "log"
//    "io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the crapiest URLShortener ever!")
}

func RedirectLink(w http.ResponseWriter, r *http.Request) {
    longURL := StorageLookup(strings.Trim(r.URL.Path,"/"))
    if longURL == ""{
        http.Error(w, "URL, not found! :S", http.StatusNotFound)
    } else {
        w.Header().Set("Location", longURL)
        w.WriteHeader(301)    
    }
}

func LinkStats(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to LinkStats2!")
}

func NewLink(w http.ResponseWriter, r *http.Request) {
    log.Println("New link required")
    var shortID string
    decoder := json.NewDecoder(r.Body)
    var myRequest shortURL_struct
    err := decoder.Decode(&myRequest)
    if err != nil {
        log.Println(err)
        http.Error(w, "UPS, something went wrong! :S", http.StatusInternalServerError)
    } else {
        fmt.Fprintln(w,"TESTE->"+myRequest.Longurl+"<-")
        shortID=generateShortID()
        fmt.Fprintln(w,"TESTE2->"+shortID+"<-")
        StorageSave(shortID,myRequest.Longurl)
    }
}

func RedirectLink2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome %s!",r.URL.Path)
    // w.Header().Set("Location", "http://google.es")
    // w.WriteHeader(301)
}