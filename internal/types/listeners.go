package types

import "github.com/DisgoOrg/disgo/core"

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

func (l *Listeners) OnEvent(event core.Event) {
	for _, listener := range l.listeners {
		listener.OnEvent(l.bot, event)
	}
}
