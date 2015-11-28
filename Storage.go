package main

import (

)

//In memory storage
var inMemoryStorage map[string]string

func StorageLookup(shortID string) string {
    longURL:=inMemoryStorage[shortID]
    return longURL
}

func StorageSave(shortID string, longURL string) {
    inMemoryStorage[shortID]=longURL
    return
}

func StorageInitialization() {
    inMemoryStorage=make(map[string]string)
}