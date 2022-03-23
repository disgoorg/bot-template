package types

import (
	"github.com/disgoorg/disgo/bot"
)

func (b *Bot) LoadModules(modules []Module) {
	b.Logger.Info("Loading modules...")
	commands := NewCommandMap(b)
	listeners := NewListeners(b)

	for _, module := range modules {
		if mod, ok := module.(CommandsModule); ok {
			commands.AddAll(mod.Commands())
		}

		if mod, ok := module.(ListenerModule); ok {
			listeners.AddListener(mod)
		}
	}

	b.Logger.Infof("Loaded %d modules", len(modules))
	b.Commands = commands
	b.Listeners = listeners
}

type Module interface{}

type CommandsModule interface {
	Commands() []Command
}

type ListenerModule interface {
	OnEvent(b *Bot, event bot.Event)
}

type CommandModule struct {
	Cmds []Command
}

func (m CommandModule) Commands() []Command {
	return m.Cmds
}
