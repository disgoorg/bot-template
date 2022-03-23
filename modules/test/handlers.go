package test

import (
	"github.com/disgoorg/bot-template/internal/types"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func testHandler(b *types.Bot, e *events.ApplicationCommandInteractionEvent) error {
	return e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("Test command").Build())
}
