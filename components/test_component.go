package components

import (
	"github.com/disgoorg/bot-template"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/json"
	"github.com/disgoorg/handler"
)

func TestComponent(b *bot_template.Bot) handler.Component {
	return handler.Component{
		Name: "test_button",
		Handler: func(e *events.ComponentInteractionCreate) error {
			return e.UpdateMessage(discord.MessageUpdate{
				Content: json.NewPtr("This is a test button update"),
			})
		},
	}
}
