package test

import (
	"github.com/DisgoOrg/bot-template/internal/types"
	"github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
)

func testHandler(b *types.Bot, e *events.ApplicationCommandInteractionEvent) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("Test command").Build())
}
