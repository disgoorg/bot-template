package commands

import (
	"github.com/disgoorg/bot-template/tbot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var TestCommand = tbot.Command{
	Create: discord.SlashCommandCreate{
		CommandName: "test",
		Description: "Test command",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				OptionName:   "choice",
				Description:  "some autocomplete choice",
				Required:     true,
				Autocomplete: true,
			},
		},
	},
	CommandHandlers: map[string]tbot.CommandHandler{
		"": func(b *tbot.Bot, e *events.ApplicationCommandInteractionCreate) error {
			return e.CreateMessage(discord.NewMessageCreateBuilder().
				SetContentf("Test command. Choice: %s", e.SlashCommandInteractionData().String("choice")).
				AddActionRow(discord.NewPrimaryButton("Test", "test_button")).
				Build(),
			)
		},
	},
	AutocompleteHandlers: map[string]tbot.AutocompleteHandler{
		"": func(b *tbot.Bot, e *events.AutocompleteInteractionCreate) error {
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
