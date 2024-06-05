package commands

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var test = discord.SlashCommandCreate{
	Name:        "test",
	Description: "test command",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:         "choice",
			Description:  "some autocomplete choice",
			Required:     true,
			Autocomplete: true,
		},
	},
}

func TestHandler(e *handler.CommandEvent) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().
		SetContentf("test command. Choice: %s", e.SlashCommandInteractionData().String("choice")).
		AddActionRow(discord.NewPrimaryButton("test", "/test-button")).
		Build(),
	)
}

func TestAutocompleteHandler(e *handler.AutocompleteEvent) error {
	return e.AutocompleteResult([]discord.AutocompleteChoice{
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
}
