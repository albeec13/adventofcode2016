package main

import (
    "fmt"
    "strings"
    "strconv"
    "regexp"
    "log"
)

/* Generic Chip Carrier type */
type ChipCarrier struct {
    chips   []int
    getChip chan int
    done    chan string
    ready   bool
    id      int
}

/* Generic Chip Carrier init routine:
   Prepare chip slice to hold 2 chips, configure channel for 
   receiving chips, and set callback channel to 'done' for sending
   status when the chipcarrier completes it's mission */
func (c *ChipCarrier) Init(done chan string, id int) {
    c.chips = make([]int, 0)
    c.getChip = make(chan int, 2)
    c.done = done
    c.id = id
    c.ready = true
}

/* Output type, which doesn't send anything, only receives */
type Output struct {
    ChipCarrier
}

func (o *Output) Init(done chan string, id int) {
    o.ChipCarrier.Init(done, id)
}

/* Processing routine, loops forever, waiting for incoming chips, until
   it has received 2 chips, then shuts itself down and reports back */
func (o *Output) Process() {
Done:
    for {
        select {
        case getChip := <-o.getChip:
            if len(o.chips) < 2 {
                o.chips = append(o.chips, getChip)

                if len(o.chips) == 2 {
                    o.done <- fmt.Sprintln("Output",o.id,"received chips:",o.chips[0],"and",o.chips[1])
                    break Done
                } else {
                    o.done <- fmt.Sprintln("Output",o.id,"received chip:",getChip)
                }
            }
        }
    }
}

/* Bot type, which is based on a standard ChipCarrier, but also adds
   channels to configure a high and low receiver, and pointers to
   those receivers, so that it can pass along chips once it has a 
   set of 2 */
type Bot struct {
    ChipCarrier
    high    *ChipCarrier
    low     *ChipCarrier
    setHigh chan *ChipCarrier
    setLow  chan *ChipCarrier
}

/* Create channels to receive target instructions for who to send chips
   to.  Also create high and low pointers to those targets. Proceed with
   standard init afterward */
func (b *Bot) Init(done chan string, id int) {
    b.setHigh = make(chan *ChipCarrier, 1)
    b.setLow  = make(chan *ChipCarrier, 1)
    b.high = nil
    b.low  = nil
    b.ChipCarrier.Init(done, id)
}

/* Bot wait for incoming chips or incoming target instructions, updating
   it's internal chip slice and targetd pointers as needed.  After each
   channel input is received, the Bot checks if it has all targets and
   chips it can hold, and if so, sends chips to those targets, shuts
   itself down, and reports back */
func (b *Bot) Process() {
Done:
    for {
        select {
        case getChip := <-b.getChip:
            if len(b.chips) < 2 {
                b.chips = append(b.chips, getChip)
                if b.trySend() { break Done }
            }
        case setHigh := <-b.setHigh:
            b.high = setHigh
            if b.trySend() { break Done }
        case setLow := <-b.setLow:
            b.low = setLow
            if b.trySend() { break Done }
        }
    }
}

/* Bot routine to check if it is ready to transmit chips */
func (b *Bot) trySend() bool {
    if len(b.chips) == 2 {
        if b.high != nil && b.low != nil {
            for !b.low.ready || !b.high.ready { }
            b.low.getChip <- Min(b.chips[0], b.chips[1])
            b.high.getChip <- Max(b.chips[0], b.chips[1])
            b.done <- fmt.Sprintln("Bot",b.id,"sent chips:",b.chips[0],"and",b.chips[1])
            return true
        }
    }
    return false
}

/* Main Code Entry Point */
func main() {
    /* done is the receiver channel for the main routine, where it receives feedback from bots and outputs */
    /* bots and outputs are lists of bots and outputs that appear in the instructions, expanded as needed */
    done := make(chan string, 500)
    bots := make([]Bot,0)
    outputs := make([]Output,0)

    /* rv and rv are regex strings for parsing incoming commands for incoming chip values or bot target instructions, respectively */
    rv := regexp.MustCompile(`value (?P<Value>\d+) goes to bot (?P<Bot>\d+)`)
    rb := regexp.MustCompile(`bot (?P<Bot>\d+) gives low to (?P<lowTarget>bot|output) (?P<Low>\d+) and high to (?P<highTarget>bot|output) (?P<High>\d+)`)

    /* Begin parsing instruction list, found in input.go */
    for _,s := range strings.Split(input, "\n") {
        if v := rv.FindStringSubmatch(s); v != nil {
            /* Parse incoming value commands */
            if len(v) == 3 {
                if val, err := strconv.Atoi(v[1]); err == nil {
                    if bot, err := strconv.Atoi(v[2]); err == nil {
                        /* If command was properly formatted, expand bot list as needed, and send chip */
                        bots = ExpandBotSlice(bots, bot, done)
                        bots[bot].getChip <- val
                    } else {
                        log.Fatal("Invalid bot index:",bot)
                    }
                } else {
                    log.Fatal("Invalid chip value:",val)
                }
            } else {
                log.Fatal("Invalid command string:", s)
            }
        } else if b := rb.FindStringSubmatch(s); b != nil {
            /* Parse target commands */
            if len(b) == 6 {
                if bot, err := strconv.Atoi(b[1]); err == nil {
                    if low, err := strconv.Atoi(b[3]); err == nil {
                        if high, err := strconv.Atoi(b[5]); err == nil {
                            lowTarget, highTarget := b[2], b[4]

                            /* If command was properly formatted, expand bot and output lists as needed and 
                               configure low and high targets on desired bot */
                            bots = ExpandBotSlice(bots, bot, done)

                            switch lowTarget {
                            case "bot":
                                bots = ExpandBotSlice(bots, low, done)
                                bots[bot].setLow <- &(bots[low].ChipCarrier)
                            case "output":
                                outputs = ExpandOutputSlice(outputs, low, done)
                                bots[bot].setLow <- &(outputs[low].ChipCarrier)
                            default:
                                log.Fatal("Invalid chip target:",lowTarget)
                            }

                            switch highTarget {
                            case "bot":
                                bots = ExpandBotSlice(bots, high, done)
                                bots[bot].setHigh <- &(bots[high].ChipCarrier)
                            case "output":
                                outputs = ExpandOutputSlice(outputs, high, done)
                                bots[bot].setHigh <- &(outputs[high].ChipCarrier)
                            default:
                                log.Fatal("Invalid chip targetL",highTarget)
                            }
                        } else {
                            log.Fatal("Invalid high target index:",high)
                        }
                    } else {
                        log.Fatal("Invalid low target index:",low)
                    }
                } else {
                    log.Fatal("Invalid bot index:",bot)
                }
            } else {
                log.Fatal("Invalid command string:",s)
            }
        } else {
            log.Fatal("Unknown command string:", s)
        }
    }

    puzz1 := ""
    puzz2 := make([]string, 0)

Complete:
    for{
        select {
        case ended := <-done:
            //fmt.Printf(ended)
            if strings.Contains(ended, "sent chips: 17 and 61") {
                puzz1 = ended
            }

            if strings.Contains(ended, "Output 0 ") || strings.Contains(ended, "Output 1 ") || strings.Contains(ended, "Output 2 ") {
                puzz2 = append(puzz2, ended)
            }

            if len(puzz1) > 1 && len(puzz2) == 3 {
                break Complete
            }
        default:
        }
    }

    fmt.Println()
    fmt.Printf("Puzzle 1: ")
    fmt.Printf(puzz1)
    fmt.Println()
    fmt.Println("Puzzle 2: ")
    for _,s := range puzz2 {
        fmt.Printf("\t")
        fmt.Printf(s)
    }
}

func ExpandBotSlice(bots []Bot, bot int, done chan string) []Bot {
    if len(bots) < bot + 1 {
        oldLen := len(bots)
        more := bot + 1 - oldLen

        bots = append(bots, make([]Bot, more)...)

        for i := oldLen; i < bot + 1; i++{
            bots[i] = Bot{}
            bots[i].Init(done, i)
            go bots[i].Process()
        }
    }
    for !bots[bot].ready { }
    return bots
}

func ExpandOutputSlice(outputs []Output, output int, done chan string) []Output{
    if len(outputs) < output + 1 {
        oldLen := len(outputs)
        more := output + 1 - oldLen

        outputs = append(outputs, make([]Output, more)...)

        for i := oldLen; i < output + 1; i++{
            outputs[i] = Output{}
            outputs[i].Init(done, i)
            go outputs[i].Process()
        }
    }
    for !outputs[output].ready { }
    return outputs
}


func Min(x int, y int) int {
    if x > y { return y } else { return x }
}

func Max(x int, y int) int {
    if x < y { return y } else { return x }
}

