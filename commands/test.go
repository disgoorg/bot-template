package commands

import (
	"github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
	"github.com/your-name/your-repo/internal/types"
)

var testCommand = types.Command{
	Create: discord.SlashCommandCreate{
		Name:              "test",
		Description:       "Test command",
		DefaultPermission: true,
	},
	Handler: func(b *types.Bot, e *events.ApplicationCommandInteractionEvent) error {
		return e.CreateMessage(discord.NewMessageCreateBuilder().SetContent("Test command").Build())
	},
}
