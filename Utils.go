package main

import (
    "math/rand"
    "time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateShortID() string {
    for shortIDLength := 5; shortIDLength < 10; shortIDLength++ {
        for i := 0; i < 6; i++ {
            shortURL:=randString(shortIDLength)
            if StorageLookup(shortURL) == "" {
                return shortURL
            }
        }
    }
    return ""
}

func randString(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func Initializations(){
    StorageInitialization()
    StatInitialization()
    rand.Seed( time.Now().UTC().UnixNano())
}