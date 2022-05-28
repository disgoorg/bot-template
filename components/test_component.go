package components

import (
	"github.com/disgoorg/bot-template/tbot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/json"
)

var TestComponent = tbot.Component{
	Action: "test_button",
	Handler: func(b *tbot.Bot, data []string, e *events.ComponentInteractionCreate) error {
		return e.UpdateMessage(discord.MessageUpdate{
			Content: json.NewPtr("This is a test button update"),
		})
	},
}
