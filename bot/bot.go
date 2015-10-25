package bot

import (
    "fmt"
    "net"
    "goda/modules"
)

type Bot struct {
    server string
    port string
    user string
    conn net.Conn
    Channel string
    Modules *modules.ModuleData
    Nick string
}

func NewBot(activeModules []string) *Bot {
    return &Bot {
        server: "lem0n.net",
        port: "6667",
        conn: nil,
        user: "noob",
        Channel: "#beepboop",
        Modules: modules.NewModuleData(activeModules),
        Nick: "notgouda",
    }
}

func (bot *Bot) Connect() (conn net.Conn, err error) {
    conn, err = net.Dial("tcp", bot.server + ":" + bot.port)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    bot.conn = conn
    fmt.Printf("Connected\n")
    return bot.conn, nil
}

func (bot *Bot) Write(text string) {
    fmt.Fprintf(bot.conn, "PRIVMSG %s :%s\r\n", bot.Channel, text)
}
