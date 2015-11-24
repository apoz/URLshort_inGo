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

func StorageInitialValues() {
    inMemoryStorage=make(map[string]string)
    inMemoryStorage["1234"]="http://google.es"
    inMemoryStorage["asdf"]="http://ddg.com"
    inMemoryStorage["qwert"]="http://bing.com"
}