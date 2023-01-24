package handlers

import (
	"github.com/disgoorg/bot-template"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
)

func MessageHandler(b *dbot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.MessageCreate) {
		// TODO: handle message
	})
}
