package bot

import (
    "bufio"
    "net"
    "net/textproto"
    "strings"

    "goda/modules"
)

type Parser struct {
    bot *Bot
    conn net.Conn
}

func NewParser(_bot *Bot) *Parser {
    return &Parser {
        bot: _bot,
        conn: _bot.conn,
    }
}

func tokenise(line string) (string, string) {
    nick := strings.Split(line, "!")[0][1:]
    message := strings.Split(line, " ")[3:]
    message[0] = message[0][1:]
    return nick, strings.Join(message, " ")
}

func (parser *Parser) Parse() {
    reader := bufio.NewReader(parser.conn)
    tp := textproto.NewReader(reader)
    for {
        line, err := tp.ReadLine()
        if err != nil {
            break
        }
        if strings.Contains(line, "PRIVMSG") {
            nick, message := tokenise(line)
            modules.Hook(nick, message)
        }
    }
}
