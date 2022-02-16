package test

import (
	"github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
	"github.com/YourName/YourRepo/internal/types"
)

func testHandler(b *types.Bot, e *events.ApplicationCommandInteractionEvent) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("Test command").Build())
}
