package main

import (
    "log"
    "net/http"
    "time"
    "math/rand"
)

func main() {

    StorageInitialValues()
    rand.Seed( time.Now().UTC().UnixNano())
    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}