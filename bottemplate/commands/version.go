package commands

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"

	"github.com/disgoorg/bot-template/bottemplate"
)

var version = discord.SlashCommandCreate{
	Name:        "version",
	Description: "version command",
}

func VersionHandler(b *bottemplate.Bot) handler.CommandHandler {
	return func(e *handler.CommandEvent) error {
		return e.CreateMessage(discord.MessageCreate{
			Content: fmt.Sprintf("Version: %s\nCommit: %s", b.Version, b.Commit),
		})
	}
}
