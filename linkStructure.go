package main

type shortURL_struct struct {
    Longurl string
}

type newURL_struct struct {
    ShortID string
    LongURL string
}

type URLstats_struct struct {
    ShortID string
    NumberOfVisits int
}