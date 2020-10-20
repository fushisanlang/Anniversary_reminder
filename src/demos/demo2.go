package main

import (
    "flag"
    "fmt"
)

var cli = flag.String("cli", "off", "cli on/off")

func main() {
    flag.Parse()
    fmt.Println(*cli)
}
