package main

import (
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func parseCommand(text string, entities []tb.MessageEntity) (string, []string) {
	text = strings.TrimSpace(strings.ToLower(text))
	var command string = ""
	var args []string = []string{}
	for _, entity := range entities {
		if entity.Type == "bot_command" {
			command_entity := text[entity.Offset : entity.Offset+entity.Length]
			command = command_entity[1:]
			args = strings.Split(strings.ReplaceAll(text, command_entity, ""), " ")
		}
	}

	return command, args
}
