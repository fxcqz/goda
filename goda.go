package main

import (
    "fmt"
    "goda/bot"
)

func main() {
    activeModules := []string{"crud"}
    goda := bot.NewBot(activeModules)
    goda.RegisterModules()
    conn, _ := goda.Connect()

    fmt.Fprintf(conn, "USER %s 8 * :%s\r\n", goda.Nick, goda.Nick)
    fmt.Fprintf(conn, "NICK %s\r\n", goda.Nick)
    fmt.Fprintf(conn, "JOIN %s\r\n", goda.Channel)
    defer conn.Close()

    parser := bot.NewParser(goda)
    parser.Parse()
}
