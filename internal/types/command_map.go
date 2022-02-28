package types

import (
	"strings"

	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/core/events"
)

func NewCommandMap(bot *Bot) *CommandMap {
	return &CommandMap{
		bot:      bot,
		commands: make(map[string]Command),
	}
}

type CommandMap struct {
	bot      *Bot
	commands map[string]Command
}

func (m *CommandMap) OnEvent(event core.Event) {
	if e, ok := event.(*events.ApplicationCommandInteractionEvent); ok {
		if cmd, ok := m.commands[e.Data.Name()]; ok {
			switch d := e.Data.(type) {
			case core.SlashCommandInteractionData:
				var name string
				if d.SubCommandGroupName != nil && d.SubCommandName != nil {
					name = *d.SubCommandGroupName + "/" + *d.SubCommandName
				} else if d.SubCommandName != nil {
					name = *d.SubCommandName
				}
				if cmd.CommandHandler == nil {
					m.bot.Logger.Errorf("No command handler for \"%s\"", e.Data.Name())
					return
				}
				err := cmd.CommandHandler[buildCommandPath(d.SubCommandName, d.SubCommandGroupName)](m.bot, e)
				if err != nil {
					m.bot.Logger.Errorf("Failed to handle command \"%s\": %s", name, err)

				}
			}
		}
	} else if e, ok := event.(*events.AutocompleteInteractionEvent); ok {
		if cmd, ok := m.commands[e.Data.CommandName]; ok {
			if cmd.AutoCompleteHandler == nil {
				m.bot.Logger.Errorf("No autocomplete handler for command \"%s\"", e.Data.CommandName)
				return
			}
			if err := cmd.AutoCompleteHandler[buildCommandPath(e.Data.SubCommandName, e.Data.SubCommandGroupName)](m.bot, e); err != nil {
				m.bot.Logger.Errorf("Failed to handle autocomplete for \"%s\": %s", e.Data.CommandName, err)
			}
		}
	} else if e, ok := event.(*events.ComponentInteractionEvent); ok {
		customID := e.Data.ID().String()
		if !strings.HasPrefix(customID, "cmd:") {
			return
		}
		args := strings.Split(customID, ":")
		cmdHandler, action := args[1], args[2]
		cmdName := strings.Split(cmdHandler, "/")[0]
		if cmd, ok := m.commands[cmdName]; ok {
			if cmd.ComponentHandler == nil {
				m.bot.Logger.Errorf("No component handler for command \"%s\"", cmdName)
				return
			}
			if h, ok := cmd.ComponentHandler[cmdHandler]; ok {
				if err := h(m.bot, e, action); err != nil {
					m.bot.Logger.Errorf("Failed to handle component interaction for \"%s\": %s", cmdName, err)
				}
				return
			}
			m.bot.Logger.Errorf("No component handler for action \"%s\" on command \"%s\"", action, cmdName)
		}
	}
}

func buildCommandPath(subcommand *string, subcommandGroup *string) string {
	var path string
	if subcommand != nil {
		path = *subcommand
	}
	if subcommandGroup != nil {
		path += "/" + *subcommandGroup
	}
	return path
}

func (m *CommandMap) AddAll(c []Command) {
	for _, cmd := range c {
		m.commands[cmd.Create.Name()] = cmd
	}
}
