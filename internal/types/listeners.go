package types

import "github.com/disgoorg/disgo/bot"

func NewListeners(b *Bot) *Listeners {
	return &Listeners{
		bot: b,
	}
}

type Listeners struct {
	bot       *Bot
	listeners []ListenerModule
}

func (l *Listeners) AddListener(listener ListenerModule) {
	l.listeners = append(l.listeners, listener)
}

func (l *Listeners) OnEvent(event bot.Event) {
	for _, listener := range l.listeners {
		listener.OnEvent(l.bot, event)
	}
}
