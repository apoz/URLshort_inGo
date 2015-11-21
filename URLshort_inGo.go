package main

import (
    "log"
    "net/http"
)

func main() {

    StorageInitialValues()
    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}