package main

import (
    "fmt"
    "strings"
    "time"

    "github.com/nosixtools/solarlunar"
)

func main() {

    timeStr := time.Now().Format("2006-01-02")
    fmt.Println(timeStr)
    timestring := solarlunar.SolarToSimpleLuanr(timeStr)
    old1 := `Y`
    old2 := `M`
    old3 := `D`
    new1 := ""
    a := strings.Replace(timestring, old1, new1, -1)
    b := strings.Replace(a, old2, new1, -1)
    c := strings.Replace(b, old3, new1, -1)
    fmt.Println(c)
    timeStr2 := time.Now().Format("0102")
    fmt.Println(timeStr2)
}
