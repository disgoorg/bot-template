package handlers

import (
	"github.com/disgoorg/bot-template/bottemplate"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
)

func MessageHandler(b *bottemplate.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.MessageCreate) {
		// TODO: handle message
	})
}
