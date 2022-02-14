package types

import (
	"strings"

	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/core/events"
)

type CommandMap struct {
	benny    *Bot
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
				if name != "" {
					if h := cmd.SubCommandHandler[name]; h != nil {
						if err := h(m.benny, e); err != nil {
							m.benny.Logger.Errorf("Failed to handle subcommand \"%s\": %s", name, err)
						}

						return
					}
					m.benny.Logger.Errorf("No subcommand handler for \"%s\" on command \"%s\"", name, e.Data.Name())
					return
				}
				if h := cmd.Handler; h != nil {
					err := h(m.benny, e)
					if err != nil {
						m.benny.Logger.Errorf("Failed to handle command \"%s\": %s", name, err)

					}
					return
				}
				m.benny.Logger.Errorf("No command handler for \"%s\"", e.Data.Name())
			}
		}
	} else if e, ok := event.(*events.AutocompleteInteractionEvent); ok {
		if cmd, ok := m.commands[e.Data.CommandName]; ok {
			if cmd.AutoCompleteHandler != nil {
				if err := cmd.AutoCompleteHandler(m.benny, e); err != nil {
					m.benny.Logger.Errorf("Failed to handle autocomplete for \"%s\": %s", e.Data.CommandName, err)
				}
				return
			}
			m.benny.Logger.Errorf("No autocomplete handler for command \"%s\"", e.Data.CommandName)
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
				m.benny.Logger.Errorf("No component handler for command \"%s\"", cmdName)
				return
			}
			if h, ok := cmd.ComponentHandler[cmdHandler]; ok {
				if err := h(m.benny, e, action); err != nil {
					m.benny.Logger.Errorf("Failed to handle component interaction for \"%s\": %s", cmdName, err)
				}
				return
			}
			m.benny.Logger.Errorf("No component handler for action \"%s\" on command \"%s\"", action, cmdName)
		}
	}
}

func (m *CommandMap) SetCommand(c Command) {
	m.commands[c.Create.Name] = c
}