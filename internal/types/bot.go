package types

import (
	"context"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
	"github.com/uptrace/bun"
)

type Bot struct {
	Client    bot.Client
	Logger    log.Logger
	Commands  *CommandMap
	Listeners *Listeners
	DB        *bun.DB
	Config    Config
	Version   string
}

func (b *Bot) SetupBot() (err error) {
	b.Client, err = disgo.New(b.Config.Token,
		bot.WithLogger(b.Logger),
		bot.WithGatewayConfigOpts(gateway.WithGatewayIntents(discord.GatewayIntentGuilds)),
		bot.WithEventListeners(b.Commands, b.Listeners),
	)
	return err
}

func (b *Bot) StartBot() (err error) {
	return b.Client.ConnectGateway(context.TODO())
}
