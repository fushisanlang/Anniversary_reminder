package main

import (
    "Anniversary_reminder/until"
    "fmt"
    "log"
    "strings"

    "github.com/robfig/cron"
)

func CronTask() {
    log.Println("********  *******  *******")
}

func CronTest() {
    cronvalue := until.ReadConf("cronvalue")
    fmt.Println(cronvalue)
    log.Println("Starting Cron...")
    c := cron.New(cron.WithSeconds())
    c.AddFunc(cronvalue, CronTask)
    c.Start()

    select {}
}

func main() {
    fmt.Println(strings.Repeat("START ", 15))
    CronTest()
    fmt.Println(strings.Repeat("END ", 15))
}
