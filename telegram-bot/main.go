package main

import (
	"fmt"
	"log"

	tb_commands "telebot/commands"

	tb "gopkg.in/tucnak/telebot.v2"
)

func FindCommand(command string) tb_commands.Command {
	for _, c := range tb_commands.GetCommands() {
		if c.Name == command {
			return c
		}
	}
	return tb_commands.Command{}
}

func main() {
	poller := getPoller()
	b, err := tb.NewBot(tb.Settings{
		Token:  config.Token,
		Poller: tb.NewMiddlewarePoller(poller, getRatelimitMiddleware),
	})
	if err != nil {
		log.Fatal(err)
	}
	b.Handle(tb.OnText, func(m *tb.Message) {
		command, args := parseCommand(m.Text, m.Entities)
		if command != "" {
			c := FindCommand(command)
			if len(c.Name) > 0 {
				if c.OwnerOnly && m.Sender.ID != config.OwnerID {
					b.Send(m.Chat, "You are not allowed to use this command.")
					return
				} else if c.PrivateOnly && m.Chat.Type != "private" {
					b.Send(m.Chat, "This command is only available in private chat.")
					return
				} else {
					go c.Execute(b, m, args)
					logMessage := fmt.Sprintf("%v execute %v on %v", m.Sender.ID, command, m.Chat.ID)
					fmt.Println(logMessage)
				}
			}
		}
	})
	b.Start()
}
