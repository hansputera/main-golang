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
			command_entity_mentionop := text[entity.Offset : entity.Offset+entity.Length]
			command_entity := strings.Split(command_entity_mentionop, "@")[0]
			command = command_entity_mentionop[1:]
			args_entity := strings.Split(strings.TrimSpace(strings.ReplaceAll(text, command_entity, "")), " ")
			for _, arg := range args_entity {
				if len(arg) > 0 {
					args = append(args, arg)
				}
			}
		}
	}

	return command, args
}
