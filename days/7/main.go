package main

import (
    "fmt"
    "strings"
    "log"
)

func main () {
    cnt := 0
    cnt2 := 0

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

        if hasABA_BAB(inBrack, outBrack) {
            cnt2++
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
    fmt.Println("ABBA: ", cnt)
    fmt.Println("ABA:  ", cnt2)
}

func hasABBA(s string) bool {
    if len(s) >= 4 {
        for i := 4; i <= len(s); i++ {
            q := s[i-4:i]
            if (q[0] == q[3]) && (q[1] == q[2]) && (q[0] != q[1]) { return true }
        }
    }
    return false
}

func hasABA_BAB(in []string, out []string) bool {
    foundIn := make([]string,0)

    for _,s := range in {
        foundIn = append(foundIn, hasABA(s)...)
    }

    for _,s := range out {
        foundOut := hasABA(s)

        for _,a := range foundOut {
            if len(a) == 3 {
                inv := []byte{' ',' ',' '}
                inv[0] = a[1]
                inv[2] = a[1]
                inv[1] = a[0]

                for _,i := range foundIn {
                    if(string(inv) == i) { return true }
                }

            }
        }
    }
    return false
}

func hasABA(s string) []string {
    found := make([]string,0)

    if len(s) >= 3 {
        for i := 3; i <= len(s); i++ {
            q := s[i-3:i]
            if (q[0] == q[2]) && (q[0] != q[1]) {
                found = append(found, q)
            }
        }
    }
    return found
}

func tern(comp bool, valTrue string, valFalse string) string {
    if comp { return valTrue } else { return valFalse }
}
