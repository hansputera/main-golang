package commands

import (
	"net/url"
	"strings"

	shortener "telebot/services"

	tb "gopkg.in/tucnak/telebot.v2"
)

func shortSIDChannel(url string, ch chan string) {
	result, err := shortener.Shortener_SDotID(url)
	if err != nil {
		ch <- "Error: " + err.Error()
		return
	} else {
		ch <- result.Url
		return
	}
}

func ShortenerSID() Command {
	return Command{
		Name:        "shortensid",
		Description: "Shorten url to s.id",
		Aliases:     []string{"sid"},
		OwnerOnly:   false,
		PrivateOnly: false,
		Execute: func(bot *tb.Bot, m *tb.Message, args []string) {
			if len(args) < 1 {
				bot.Send(m.Chat, "Usage: /shorten_sid <url>")
				return
			} else {
				url_input := args[0]
				if _, errValid := url.ParseRequestURI(url_input); errValid != nil {
					bot.Send(m.Chat, "Invalid url")
					return
				} else {
					mEdit, errMedit := bot.Send(m.Sender, "Shortening")
					if errMedit != nil {
						bot.Send(m.Chat, "Please allow me to send private message!")
						return
					} else {
						channel := make(chan string)
						go shortSIDChannel(url_input, channel)

						result := <-channel
						if strings.HasPrefix(result, "Error") {
							bot.Edit(mEdit, result)
							return
						} else {
							bot.Edit(mEdit, "Shortened result: "+result)
							return
						}
					}
				}
			}
		},
	}
}
