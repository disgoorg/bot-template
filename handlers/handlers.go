package handlers

import (
	"github.com/disgoorg/bot-template/tbot"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
)

var MessageHandler = func(b *tbot.Bot) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.MessageCreate) {

	})
}
