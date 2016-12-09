package main

import (
    "fmt"
    "strings"
    "log"
)

func main () {
    cnt := 0

Top:
    for _,in := range strings.Split(input, "\n") {
        inBrack := make([]string,0)
        outBrack := make([]string, 0)
        inner := false
        start := 0
        for {
            sepi := strings.IndexAny(in, "[]")
            if (sepi != -1) {
                str := in[start:sepi]
                if str != "" {
                    if !inner && (string(in[sepi]) == "[") {
                        outBrack = append(outBrack, str)
                    } else if inner && (string(in[sepi]) == "]") {
                        inBrack = append(inBrack, str)
                    } else {
                        log.Fatal("Invalid input.")
                    }
                }
                inner = !inner
                in = strings.TrimPrefix(in, str + tern(inner,"[","]"))
                start = 0
            } else if in == "" {
                break
            } else {
                if str := in[start:]; str != "" {
                    if inner {
                        inBrack = append(inBrack, str)
                    } else {
                        outBrack = append(outBrack, str)
                    }
                }
                break
            }
        }

        for _,str := range inBrack {
            if hasABBA(str) {
                continue Top
            }
        }

        for _,str := range outBrack {
            if hasABBA(str) {
                cnt++
                continue Top
            }
        }
    }
    fmt.Println(cnt)
}

func hasABBA(s string) bool {
    ret := false

    if len(s) >= 4 {
        for i := 4; i <= len(s); i++ {
            q := s[i-4:i]
            if (q[0] == q[3]) && (q[1] == q[2]) && (q[0] != q[1]) { return true }
        }
    }

    return ret
}

func tern(comp bool, valTrue string, valFalse string) string {
    if comp { return valTrue } else { return valFalse }
}
