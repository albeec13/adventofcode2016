package main

import (
    "fmt"
    "strings"
    "strconv"
    "log"
)

func main() {
    output := recurseDecomp(input)
    fmt.Println("Length :", output)
}

func recurseDecomp(input string) (output int) {
    for start := strings.Index(input, "("); start != -1; start = strings.Index(input, "(")  {
        end := strings.Index(input, ")");

        if (start + 1)  < (end -1) {
            marker := strings.Split(input[start+1:end], "x")
            var err error
            cnt, rep := 0, 0

            if len(marker) == 2 {
                if cnt, err = strconv.Atoi(marker[0]); err == nil {
                    if rep, err = strconv.Atoi(marker[1]); err == nil {
                        output = output + len(input[:start])

                        substr := input[end + 1 : end + 1 + cnt]

                        if strings.Index(substr, "(") != -1 {
                            output = output + (rep * recurseDecomp(substr))
                        } else {
                            output = output + (rep * cnt)
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
    output = output + len(input[:])
    return output
}
