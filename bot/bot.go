package bot

import (
    "fmt"
    "net"
    "reflect"
    "strings"
    "goda/modules"
)

type Bot struct {
    server string
    port string
    user string
    conn net.Conn
    Channel string
    Modules *modules.ModuleData
    Hooks []*modules.Module
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

func (bot *Bot) Register(module string) *modules.Module {
    module = strings.ToLower(module)
    module = string(module[0] ^ 0x20) + module[1:]
    _type, _ := modules.TypeRegistry.Get(module)
    rcmds := reflect.ValueOf(_type).MethodByName("Commands").Call([]reflect.Value{})
    commands := rcmds[0].Interface().(map[string]string)
    return modules.NewModule(module, _type, commands)
}

func (bot *Bot) RegisterModules() {
    var hooks []*modules.Module
    for _, module := range bot.Modules.Active {
        hooks = append(hooks, bot.Register(module))
    }
    bot.Hooks = hooks
}

func (parser *Parser) Hook(nick string, message string) {
    words := strings.Split(message, " ")
    if words[0][:1] == "%" {
        // calling a command
        for _, mod := range parser.bot.Hooks {
            command := words[0][1:]
            if cmd, ok := mod.Commands[command]; ok {
                parser.bot.Write(reflect.ValueOf(mod.Type).MethodByName(cmd).Call([]reflect.Value{})[0].Interface().(string))
            }
        }
    }
}
