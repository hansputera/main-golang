package commands

import (
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
)

func Help() Command {
	return Command{
		Name:        "help",
		Description: "Show help",
		Aliases:     []string{"h", "halp"},
		OwnerOnly:   false,
		PrivateOnly: true,
		Execute: func(bot *tb.Bot, m *tb.Message, args []string) {
			commands := GetCommands()
			var commands_string string = ""
			for index, command := range commands {
				commands_string += fmt.Sprintf("`%v`. %v (%v)\n", index+1, command.Name, command.Description)
			}
			commands_string = fmt.Sprintf("My Command list:\n\n%v", commands_string)
			bot.Send(m.Chat, commands_string, &tb.SendOptions{
				ParseMode: "Markdown",
			})
		},
	}
}
