package main

import (
    "fmt"
    "strings"
    "strconv"
    "log"
)

type Screen [][]bool

func main() {
    cnt := 0
    screen := make(Screen, 6, 6)

    for i,_ := range screen {
        screen[i] = make([]bool,50,50)
    }

    for _,command := range strings.Split(input, "\n") {
        screen.ProcessCmd(command)
    }

    for _,row := range screen {
        for _,val := range row {
            if val {
                cnt++
            }
        }
    }

    fmt.Println(screen)
    fmt.Println("Pixels lit: ", cnt)
}

func (p *Screen) ProcessCmd(cmd string) {
    if cmdSlice := strings.Fields(cmd); len(cmdSlice) < 2 {
        log.Fatal("Invalid input: ", cmd)
    } else {
        switch cmdSlice[0] {
        case "rotate":
            p.Rotate(cmdSlice[1:])
        case "rect":
            p.Rect(cmdSlice[1:])
        default:
            log.Fatal("Invalid command in input: ", cmd)
        }
    }
}

func (p *Screen) Rotate(cmd []string) {
    if len(cmd) != 4 {
        log.Fatal("Invalid command in input: ", strings.Join(cmd," "))
    } else {
        switch cmd[0] {
        case "row":
            p.RotateRow(cmd[1:])
        case "column":
            p.RotateColumn(cmd[1:])
        default:
            log.Fatal("Invalid command in input: ", strings.Join(cmd," "))
        }
    }
}

func (p *Screen) Rect(cmd []string) {
    if len(cmd) > 1 {
        log.Fatal("Invalid command in input: ", strings.Join(cmd," "))
    } else {
        if rect := strings.Split(cmd[0], "x"); len(rect) == 2 {
            if x, err := strconv.Atoi(rect[0]); err == nil {
                if y, err := strconv.Atoi(rect[1]); err == nil {
                    for j := 0; j < y; j++ {
                        for i := 0; i < x; i++ {
                            (*p)[j][i] = true
                        }
                    }
                } else {
                    log.Fatal("Non-integer in rect input: ", strings.Join(cmd," "))
                }
            } else {
                log.Fatal("Non-integer in rect input: ", strings.Join(cmd," "))
            }
        } else {
            log.Fatal("Invalid rect command: ", strings.Join(cmd," "))
        }
    }
}

func (p *Screen) RotateRow(cmd []string) {
    if (len(cmd) != 3) || (!strings.HasPrefix(cmd[0], "y=")) || !(cmd[1] == "by") {
        log.Fatal("Invalid command in input: ", strings.Join(cmd," "))
    } else {
        if row, err := strconv.Atoi(cmd[0][2:]); (err == nil) && (row < 7) {
            if n, err := strconv.Atoi(cmd[2]); err == nil {
                rowLen := len((*p)[row])
                n = n % rowLen
                r, l := (*p)[row][:(rowLen - n)], (*p)[row][(rowLen - n):]
                (*p)[row] = append(l,r...)
            } else {
                log.Fatal("Non-integer rotation amount in command: ", strings.Join(cmd," "))
            }
        } else {
            log.Fatal("Invalid row number in command: ", strings.Join(cmd," "))
        }
    }
}

func (p *Screen) RotateColumn(cmd []string) {
    if (len(cmd) != 3) || (!strings.HasPrefix(cmd[0], "x=")) || !(cmd[1] == "by") {
        log.Fatal("Invalid command in input: ", strings.Join(cmd," "))
    } else {
        if col, err := strconv.Atoi(cmd[0][2:]); (err == nil) && (col < 51) {
            if n, err := strconv.Atoi(cmd[2]); err == nil {
                colLen := len((*p))
                n = n % colLen

                rep := func(i int) int {
                    return ((i - n) + len((*p))) % len(*p)
                }

                temp := make([]bool, colLen)

                for i,_ := range temp {
                    temp[i] = (*p)[rep(i)][col]
                }

                for i,row := range (*p) {
                    row[col] = temp[i]
                }
            } else {
                log.Fatal("Non-integer rotation amount in command: ", strings.Join(cmd," "))
            }
        } else {
            log.Fatal("Invalid row number in command: ", strings.Join(cmd," "))
        }
    }
}

func (p Screen) String() string {
    ret := ""
    for _,row := range p {
        for _,col := range row {
            if col {
                ret = ret + "\u2588"
            } else {
                ret = ret + "\u2591"
            }
        }
        ret = ret + "\n"
    }
    return ret
}
