package main

import (
    "fmt"
    "os"
    "bufio"
    "github.com/pomke/robot"
)

func main() {
    //Create a new robot with a 5x5 table
    r := robot.NewRobot(&robot.Table {5, 5})
    c := robot.NewController(r)

    //Read lines from stdin and send them to our Controller
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        c.DoCommand(input.Text())
    }
    if err := input.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}
