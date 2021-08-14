package main

import (
	"log"
	"strings"

	cmd "telebot/commands"

	tb "gopkg.in/tucnak/telebot.v2"
)

func parseCommand(bot *tb.Bot, text string, entities []tb.MessageEntity) (string, []string) {
	username := bot.Me.Username
	text = strings.TrimSpace(strings.ToLower(text))
	var command string = ""
	var args []string = []string{}
	for _, entity := range entities {
		if entity.Type == "bot_command" {
			command_entity_mentionop := text[entity.Offset : entity.Offset+entity.Length]
			command_entities := strings.Split(command_entity_mentionop, "@")
			command_entity := ""
			if len(command_entities) > 1 {
				usernameEntity := strings.Split(command_entities[1], " ")[0]
				if usernameEntity == username {
					command_entity = command_entities[0]
				}
			} else {
				command_entity = command_entities[0]
			}
			command = command_entity[1:]
			args_entity := strings.Split(strings.TrimSpace(strings.ReplaceAll(text, command_entity_mentionop, "")), " ")
			for _, arg := range args_entity {
				if len(arg) > 0 {
					args = append(args, arg)
				}
			}
		}
	}

	return command, args
}

func registerCommands(bot *tb.Bot) {
	commands := cmd.GetCommands()
	commands_telegram := []tb.Command{}
	for _, command := range commands {
		commands_telegram = append(commands_telegram, tb.Command{
			Text:        command.Name,
			Description: command.Description,
		})
	}

	log.Println("Registering commands...")
	err := bot.SetCommands(commands_telegram)
	if err != nil {
		log.Default().Fatalln(err)
	} else {
		log.Default().Println("Commands registered.")
	}
}
