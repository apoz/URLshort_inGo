package main

import (
    "log"
    "strconv"
)

//In memory storage for statistics
var inMemoryStatistics map[string]int

func StatLookup(shortID string) int {
    numberOfVisits:=inMemoryStatistics[shortID]
    return numberOfVisits
}

func StatIncrease(shortID string) {

    if inMemoryStatistics[shortID]==0 {
        inMemoryStatistics[shortID] = 1
    } else {
        inMemoryStatistics[shortID] += 1
    }
    return
}

func StatInitialization() {
    inMemoryStatistics=make(map[string]int)
}