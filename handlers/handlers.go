package handlers

import (
	"github.com/disgoorg/bot-template"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
)

var MessageHandler = func(b *bot_template.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.MessageCreate) {

	})
}
