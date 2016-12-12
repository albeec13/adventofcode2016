package main

import (
    "fmt"
    "strings"
    "strconv"
    "log"
)

func main() {
    var output string
    //next := 0

    for start := strings.Index(input, "("); start != -1; start = strings.Index(input, "(")  {
        end := strings.Index(input, ")");

        if (start + 1)  < (end -1) {
            marker := strings.Split(input[start+1:end], "x")
            var err error
            cnt, rep := 0, 0

            if len(marker) == 2 {
                if cnt, err = strconv.Atoi(marker[0]); err == nil {
                    if rep, err = strconv.Atoi(marker[1]); err == nil {
                        output = output + input[:start]
                        for i := 0; i < rep; i++ {
                            if len(input[end+1:]) >= cnt {
                                output = output + input[end + 1 : end + 1 + cnt]
                            } else {
                                log.Fatal("Invalid input file length")
                            }
                        }
                        input = input[end + 1 + cnt:]
                    } else {
                        log.Fatal("Non-integer marker: ", marker[1])
                    }
                } else {
                    log.Fatal("Non-integer marker: ", marker[0])
                }
            } else {
                log.Fatal("Invalid marker format: ", input[start+1:end])
            }
        } else {
            log.Fatal("Invalid marker format: ", input[start+1:end])
        }
    }
    output = output + input[:]
    fmt.Println("Length of output: ", len(output))
}

