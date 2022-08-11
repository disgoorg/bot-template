package commands

import (
	bot_template "github.com/disgoorg/bot-template"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/handler"
)

func TestCommand(b *bot_template.Bot) handler.Command {
	return handler.Command{
		Create: discord.SlashCommandCreate{
			Name:        "test",
			Description: "Test command",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionString{
					Name:         "choice",
					Description:  "some autocomplete choice",
					Required:     true,
					Autocomplete: true,
				},
			},
		},
		CommandHandlers: map[string]handler.CommandHandler{
			"": func(e *events.ApplicationCommandInteractionCreate) error {
				return e.CreateMessage(discord.NewMessageCreateBuilder().
					SetContentf("Test command. Choice: %s", e.SlashCommandInteractionData().String("choice")).
					AddActionRow(discord.NewPrimaryButton("Test", "handler:test_button")).
					Build(),
				)
			},
		},
		AutocompleteHandlers: map[string]handler.AutocompleteHandler{
			"": func(e *events.AutocompleteInteractionCreate) error {
				return e.Result([]discord.AutocompleteChoice{
					discord.AutocompleteChoiceString{
						Name:  "1",
						Value: "1",
					},
					discord.AutocompleteChoiceString{
						Name:  "2",
						Value: "2",
					},
					discord.AutocompleteChoiceString{
						Name:  "3",
						Value: "3",
					},
				})
			},
		},
	}
}
