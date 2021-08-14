package commands

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func Ping() Command {
	return Command{
		Name:        "ping",
		Description: "Ping command",
		Aliases:     []string{"pong"},
		OwnerOnly:   false,
		PrivateOnly: false,
		Execute: func(bot *tb.Bot, m *tb.Message, args []string) {
			bot.Send(m.Chat, "Pong!")
		},
	}
}
