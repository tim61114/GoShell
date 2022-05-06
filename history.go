//History package implements history functions
package main

import (
    "fmt"
    "strings"
)

var HistoryList [100]string
var curIndex int
var numCommands int
var size int

func HistInit(limit uint) {
    curIndex = 0   
    numCommands = 0
    size = 0
}

func HistDestroy() {
    //Hist_list
}

func HistAdd(cmd string) {

    HistoryList[curIndex] = cmd;
    curIndex = curIndex + 1 % 100 
    numCommands++
    if (size < 100) {
        size ++
    }
    
}

func HistPrint() {
    for i := size - 1 ; i >= 0; i-- {
        temp := HistoryList[(curIndex - i + 99) % 100]
        fmt.Printf("%d %s\n", numCommands - i, temp)
    }
}

func HistSearchPrefix(prefix string) (string, bool) {

    for i := size - 1; i >= 0; i-- {
        temp := HistoryList[(curIndex - i + 99) % 100]
        if (strings.HasPrefix(temp, prefix)) {
            return temp, true
        }
    }
    return "", false
}

func HistSearchCnum(commandNumber int) (string, bool) {
    if commandNumber > numCommands || numCommands - commandNumber > 99 {
        return "", false
    } else {
        index := (curIndex - (numCommands - commandNumber) + 99) % 100
        return HistoryList[index], true
    }
}

func HistLastCnum() int {
    return numCommands
}
