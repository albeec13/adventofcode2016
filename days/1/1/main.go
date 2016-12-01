package main

import (
    "fmt"
    "log"
    "strings"
    "strconv"
)

var input = "R3, R1, R4, L4, R3, R1, R1, L3, L5, L5, L3, R1, R4, L2, L1, R3, L3, R2, R1, R1, L5, L2, L1, R2, L4, R1, L2, L4, R2, R2, L2, L4, L3, R1, R4, R3, L1, R1, L5, R4, L2, R185, L2, R4, R49, L3, L4, R5, R1, R1, L1, L1, R2, L1, L4, R4, R5, R4, L3, L5, R1, R71, L1, R1, R186, L5, L2, R5, R4, R1, L5, L2, R3, R2, R5, R5, R4, R1, R4, R2, L1, R4, L1, L4, L5, L4, R4, R5, R1, L2, L4, L1, L5, L3, L5, R2, L5, R4, L4, R3, R3, R1, R4, L1, L2, R2, L1, R4, R2, R2, R5, R2, R5, L1, R1, L4, R5, R4, R2, R4, L5, R3, R2, R5, R3, L3, L5, L4, L3, L2, L2, R3, R2, L1, L1, L5, R1, L3, R3, R4, R5, L3, L5, R1, L3, L5, L5, L2, R1, L3, L1, L3, R4, L1, R3, L2, L2, R3, R3, R4, R4, R1, L4, R1, L5"

type CompassDirection int

const (
    N CompassDirection = iota
    E
    S
    W
)

type Path struct {
    North int
    East int
    Bearing CompassDirection
}

func main() {
    // start with no distances from starting point, facing north
    path := Path{0,0,N}

    // Remove white space, and create slice splitting on , from original input
    inp := strings.Split(strings.Join(strings.Fields(input),""),",")

    for _,step := range inp {
        if len(step) >= 2 {
            dir := string(step[0])
            if mag,err := strconv.Atoi(string(step[1:])); err == nil {
                path.turn(dir)
                path.move(mag)
            } else {
                log.Fatal("Bad input")
            }
        } else {
            log.Fatal("Bad input")
        }
    }

    fmt.Println(path)
    fmt.Println("Total distance: ", abs(path.North) + abs(path.East))
}

func (p *Path) turn (direction string) {
    switch direction {
    case "R":
        p.Bearing = (p.Bearing + 1) % 4
    case "L":
        p.Bearing = (p.Bearing + 3) % 4
    default:
        log.Fatal("Invalid turn")
    }
}

func (p *Path) move (magnitude int) {
    switch p.Bearing {
    case N:
        p.North += magnitude
    case E:
        p.East += magnitude
    case S:
        p.North -= magnitude
    case W:
        p.East -= magnitude
    default:
        log.Fatal("Invalid direction")
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}
