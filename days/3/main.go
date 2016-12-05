package main

import (
   "fmt"
   "strings"
   "strconv"
   "log"
)

type Triangle struct {
    Sides [3]int
}

func main() {
    cnt := 0

    triLines := strings.Split(input, "\n")
    for _,triLine := range triLines {
        var triInts [3]int
        var err error

        triple := strings.Fields(triLine)
        if len(triple) == 3 {
            for i,str := range triple {
                if triInts[i], err = strconv.Atoi(str); err != nil {
                    log.Fatal(err)
                } 
            }
        } else {
            log.Fatal("Bad input file.")
        }
        triangle := Triangle{triInts}
        cnt += triangle.Valid()
    }
    fmt.Println("Valid Triangles: ", cnt)
}

func (t *Triangle) Valid() int {
    switch {
    case t.Sides[0] + t.Sides[1] <= t.Sides[2]:
        return 0
    case t.Sides[1] + t.Sides[2] <= t.Sides[0]:
        return 0
    case t.Sides[0] + t.Sides[2] <= t.Sides[1]:
        return 0
    default:
        return 1
    }
}
