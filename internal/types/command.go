package types

import (
	"github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
)

type (
	CommandHandler      func(b *Bot, e *events.ApplicationCommandInteractionEvent) error
	ComponentHandler    func(b *Bot, e *events.ComponentInteractionEvent, action string) error
	AutocompleteHandler func(b *Bot, e *events.AutocompleteInteractionEvent) error
)

type Command struct {
	Create              discord.ApplicationCommandCreate
	CommandHandler      map[string]CommandHandler
	ComponentHandler    map[string]ComponentHandler
	AutoCompleteHandler map[string]AutocompleteHandler
}

func (b *Bot) SyncCommands() error {
	b.Logger.Info("Syncing commands...")
	var commands []discord.ApplicationCommandCreate
	for _, cmd := range b.Commands.commands {
		commands = append(commands, cmd.Create)
	}

	if b.Config.DevMode {
		b.Logger.Info("Syncing guild commands...")
		_, err := b.Bot.SetGuildCommands(b.Config.DevGuildID, commands)
		return err
	}
	b.Logger.Infof("Syncing global commands...")
	_, err := b.Bot.SetCommands(commands)
	return err
}
