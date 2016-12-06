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
    var triLineArray [3]string
    nLines := 0

    triLines := strings.Split(input, "\n")
    for _,triLine := range triLines {

        triLineArray[nLines] = triLine
        nLines++

        if nLines == 3 {
            var triInts [3][3]int
            var err error

            for j := 0; j < 3; j++ {
                triple := strings.Fields(triLineArray[j])
                if len(triple) == 3 {
                    for i,str := range triple {
                        if triInts[i][j], err = strconv.Atoi(str); err != nil {
                            log.Fatal(err)
                        }
                    }
                } else {
                    log.Fatal("Bad input file.")
                }
            }

            for i := 0; i < 3; i++ {
                triangle := Triangle{triInts[i]}
                cnt += triangle.Valid()
            }
            nLines =0
        }
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
