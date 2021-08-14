package main

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var ratelimit_users []int = []int{}

func getPoller() tb.Poller {
	return &tb.LongPoller{Timeout: time.Second * 10}
}

func IsRatelimited(userId int) bool {
	for _, user := range ratelimit_users {
		if user == userId {
			return true
		}
	}
	return false
}

func RatelimitHandle(userId int) {
	ratelimit_users = append(ratelimit_users, userId)
	time.Sleep(config.Cooldown)
	index := len(ratelimit_users) - 1
	ratelimit_users = append(ratelimit_users[:index], ratelimit_users[index+1:]...)
}

func getRatelimitMiddleware(update *tb.Update) bool {
	if update.Message != nil {
		sender := update.Message.Sender
		if sender.ID == config.OwnerID {
			return true
		} else {
			if IsRatelimited(sender.ID) {
				return false
			} else {
				go RatelimitHandle(sender.ID)
				return true
			}
		}
	}

	return true
}
