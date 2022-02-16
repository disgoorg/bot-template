package types

import "github.com/DisgoOrg/disgo/core"

func (b *Bot) LoadModules(modules []Module) {
	b.Logger.Info("Loading modules...")
	commandsMap := NewCommandMap(b)
	listeners := NewListeners(b)

	for _, module := range modules {
		if mod, ok := module.(CommandsModule); ok {
			commandsMap.AddCommands(mod.Commands())
		}

		if mod, ok := module.(ListenerModule); ok {
			listeners.AddListener(mod)
		}
	}

	b.Logger.Infof("Loaded %d modules", len(modules))
	b.Commands = commandsMap
	b.Listeners = listeners
}

type Module interface{}

type CommandsModule interface {
	Commands() []Command
}

type ListenerModule interface {
	OnEvent(b *Bot, event core.Event)
}

type CommandModule struct {
	Cmds []Command
}

func (m CommandModule) Commands() []Command {
	return m.Cmds
}
