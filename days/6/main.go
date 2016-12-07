package main

import (
    "fmt"
    "strings"
)

func main() {
    var runeCnt []map[rune]int
    var message []byte

    inputSlice := strings.Split(input, "\n")
    runeCnt = make([]map[rune]int,len(inputSlice[0]))
    message = make([]byte,len(inputSlice[0]))

    for i,_ := range runeCnt {
        runeCnt[i] = make(map[rune]int)
    }

    for _,row := range inputSlice {
        for i,char := range row {
            runeCnt[i][char]++
        }
    }

    for i,column := range runeCnt {
        min := len(inputSlice)
        letter := 'a' - 1
        for key,val := range column {
            if val < min {
                letter = key
                min = val
            }
        }
        message[i] = byte(letter)
    }

    fmt.Println(string(message))
}
