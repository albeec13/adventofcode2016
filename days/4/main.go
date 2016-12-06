package main

import (
    "fmt"
    "strings"
    "sort"
    "strconv"
    "log"
)

type RuneCnt struct {
    Char rune
    Count int
}

type ByCount []RuneCnt
func (a ByCount) Len() int		{ return len(a) }
func (a ByCount) Swap(i, j int)		{ a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool {
    switch {
    case a[i].Count == a[j].Count:
        return a[i].Char > a[j].Char
    default:
        return a[i].Count < a[j].Count
    }
}

func main() {
    cnt := 0
    inputRooms := strings.Split(input, "\n")

    for _,room := range inputRooms {
        cnt += IsRealRoom(room)
    }

    fmt.Println(cnt)
}

func IsRealRoom(roomEncrypted string) int {
    strSplt := strings.Split(roomEncrypted,"-")
    name := strings.Join(strSplt[0:len(strSplt)-1],"")
    strSplt = strings.Split(strSplt[len(strSplt)-1],"[")
    secID := strSplt[0]
    chk := strings.Split(strSplt[1],"]")[0]

    runeMap := make(map[rune]int)

    for _,char := range name {
        if runeMap[char] == 0 {
            runeMap[char] = strings.Count(name, string(char))
        }
    }

    runeCnt := make([]RuneCnt, len(runeMap))
    i := 0

    for key,val := range runeMap {
        runeCnt[i] = RuneCnt{key, val}
        i++
    }

    sort.Sort(sort.Reverse(ByCount(runeCnt)))

    if len(runeCnt) >= len([]byte(chk)) {
        for i := 0; i < len([]byte(chk)); i++ {
            if runeCnt[i].Char != rune(chk[i]) {
                return 0
            }
        }
    } else {
        return 0
    }

    if ret,err := strconv.Atoi(secID); err == nil {
        return ret
    } else {
        log.Fatal("Non-numerical sector ID")
        return 0
    }

    return 0
}

