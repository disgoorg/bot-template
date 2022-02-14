package types

import (
	"context"

	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/core/bot"
	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/disgo/gateway"
	"github.com/DisgoOrg/log"
	"github.com/uptrace/bun"
)

type Bot struct {
	Bot      *core.Bot
	Logger   log.Logger
	Commands *CommandMap
	DB       *bun.DB
	Config   Config
	Version  string
}

func (b *Bot) SetupBot() (err error) {
	b.Bot, err = bot.New(b.Config.Token,
		bot.WithLogger(b.Logger),
		bot.WithGatewayOpts(gateway.WithGatewayIntents(discord.GatewayIntentGuilds)),
		bot.WithEventListeners(b.Commands),
	)
	return err
}

func (b *Bot) StartBot() (err error) {
	return b.Bot.ConnectGateway(context.TODO())
}
