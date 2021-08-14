package events

import (
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

func EventAddedToGroup(bot *tb.Bot, m *tb.Message) {
	if !m.Private() {
		log.Default().Printf("Added to group: %v", m.Chat.ID)
		_, errMAdded := bot.Send(m.Chat, "Thanks for add me to this group :)")
		if errMAdded != nil {
			log.Default().Fatalln(errMAdded)
		}
	}
}
