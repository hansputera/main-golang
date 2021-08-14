package main

import (
	"fmt"
	"log"

	tb_commands "telebot/commands"

	events "telebot/events"

	tb "gopkg.in/tucnak/telebot.v2"
)

func FindCommand(command string) tb_commands.Command {
	for _, c := range tb_commands.GetCommands() {
		if c.Name == command {
			return c
		} else {
			for _, alias := range c.Aliases {
				if alias == command {
					return c
				}
			}
		}
	}
	return tb_commands.Command{}
}

func main() {
	for _, cmd := range tb_commands.GetCommands() {
		logCmdText := fmt.Sprintf("%v loaded with %v aliases", cmd.Name, len(cmd.Aliases))
		log.Default().Println(logCmdText)
	}
	poller := getPoller()
	b, err := tb.NewBot(tb.Settings{
		Token:  Token,
		Poller: tb.NewMiddlewarePoller(poller, getRatelimitMiddleware),
	})
	if err != nil {
		log.Fatal(err)
	}
	registerCommands(b)
	b.Handle(tb.OnText, func(m *tb.Message) {
		command, args := parseCommand(b, m.Text, m.Entities)
		if command != "" {
			c := FindCommand(command)
			if len(c.Name) > 0 {
				if c.OwnerOnly && m.Sender.ID != OwnerID {
					b.Send(m.Chat, "You are not allowed to use this command.")
					return
				} else if c.PrivateOnly && m.Chat.Type != "private" {
					b.Send(m.Chat, "This command is only available in private chat.")
					return
				} else {
					go c.Execute(b, m, args)
					logMessage := fmt.Sprintf("%v execute %v on %v", m.Sender.ID, c.Name, m.Chat.ID)
					log.Default().Println(logMessage)
				}
			}
		}
	})

	// =============== EVENTS ================= //
	b.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		events.EventAddedToGroup(b, m)
	})
	// ======================================== //
	b.Start()
}
