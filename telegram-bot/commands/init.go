package commands

import tb "gopkg.in/tucnak/telebot.v2"

type CommandRunner func(bot *tb.Bot, m *tb.Message, args []string)
type Command struct {
	Name        string
	Description string
	Aliases     []string
	Execute     CommandRunner
	OwnerOnly   bool
	PrivateOnly bool
}

func GetCommands() []Command {
	return []Command{
		Ping(),
		ShortenerSID(),
		Help(),
	}
}
