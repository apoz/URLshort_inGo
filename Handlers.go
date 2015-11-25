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
        shortID=generateShortID()
        StorageSave(shortID,myRequest.Longurl)
        
        mynewURL :=newURL_struct{shortID,myRequest.Longurl}
        js, err := json.Marshal(mynewURL)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(js)
        
    }
}

func RedirectLink2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome %s!",r.URL.Path)
    // w.Header().Set("Location", "http://google.es")
    // w.WriteHeader(301)
}