package main

import (
    "fmt"
    "goda/bot"
)

func main() {
    activeModules := []string{"core"}
    goda := bot.NewBot(activeModules)
    conn, _ := goda.Connect()

    fmt.Fprintf(conn, "USER %s 8 * :%s\r\n", goda.Nick, goda.Nick)
    fmt.Fprintf(conn, "NICK %s\r\n", goda.Nick)
    fmt.Fprintf(conn, "JOIN %s\r\n", goda.Channel)
    defer conn.Close()

    parser := bot.NewParser(goda)
    parser.Parse()
}
