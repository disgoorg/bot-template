package components

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/json"
)

func TestComponent(e *handler.ComponentEvent) error {
	return e.UpdateMessage(discord.MessageUpdate{
		Content: json.Ptr("This is a test button update"),
	})
}
