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

type Point struct {
    North int
    East int
}

type Path struct {
    Location Point
    Bearing CompassDirection
    History map[Point]int
}

func main() {
    // start with no distances from starting point, facing north
    var path Path
    path.Init(0,0,N)

    // Remove white space, and create slice splitting on , from original input
    inp := strings.Split(strings.Join(strings.Fields(input),""),",")

    for _,step := range inp {
        if len(step) >= 2 {
            dir := string(step[0])
            if mag,err := strconv.Atoi(string(step[1:])); err == nil {
                path.turn(dir)
                if done := path.move(mag); done {
                    break;
                }
            } else {
                log.Fatal("Bad input")
            }
        } else {
            log.Fatal("Bad input")
        }
    }

    fmt.Println(path.Location)
    fmt.Println("Total distance: ", abs(path.Location.North) + abs(path.Location.East))
}

func (p *Path) Init (north int, east int, bearing CompassDirection) {
    p.Location = Point{north, east}
    p.Bearing = bearing
    p.History = make(map[Point]int)
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

func (p *Path) move (magnitude int) bool {
    for i := 0; i < abs(magnitude); i++ {
        switch p.Bearing {
        case N:
            p.Location.North += 1
        case E:
            p.Location.East += 1
        case S:
            p.Location.North -= 1
        case W:
            p.Location.East -= 1
        default:
            log.Fatal("Invalid direction")
        }

        p.History[p.Location] += 1

        if p.History[p.Location] == 2 {
            fmt.Println(p.Location)
            return true
        }
    }
    return false
}

func abs(x int) int {
    if x < 0 {
        return -x
    } else {
        return x
    }
}
