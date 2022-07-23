package bot_template

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/handler"
	"github.com/disgoorg/log"
	"github.com/disgoorg/utils/paginator"
)

func New(logger log.Logger, version string, config Config) *Bot {
	return &Bot{
		Logger:    logger,
		Config:    config,
		Paginator: paginator.NewManager(),
		Version:   version,
	}
}

type Bot struct {
	Logger    log.Logger
	Handler   *handler.Handler
	Client    bot.Client
	Paginator *paginator.Manager
	Config    Config
	Version   string
}

func (b *Bot) SetupBot(listeners ...bot.EventListener) {
	b.Handler = handler.New(b.Logger)
	var err error
	b.Client, err = disgo.New(b.Config.Token,
		bot.WithLogger(b.Logger),
		bot.WithGatewayConfigOpts(gateway.WithIntents(gateway.IntentGuilds, gateway.IntentGuildMessages, gateway.IntentMessageContent)),
		bot.WithCacheConfigOpts(cache.WithCacheFlags(cache.FlagGuilds)),
		bot.WithEventListenerFunc(b.OnReady),
		bot.WithEventListeners(b.Paginator, b.Handler),
		bot.WithEventListeners(listeners...),
	)
	if err != nil {
		b.Logger.Fatal("Failed to setup b: ", err)
	}
}

func (b *Bot) StartAndBlock() {
	if err := b.Client.OpenGateway(context.TODO()); err != nil {
		b.Logger.Errorf("Failed to connect to gateway: %s", err)
	}

	b.Logger.Info("Client is running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
	b.Logger.Info("Shutting down...")
}

func (b *Bot) OnReady(_ *events.Ready) {
	b.Logger.Infof("Butler ready")
	if err := b.Client.SetPresence(context.TODO(), gateway.MessageDataPresenceUpdate{
		Activities: []discord.Activity{
			{
				Name: "you",
				Type: discord.ActivityTypeListening,
			},
		},
		Status: discord.OnlineStatusOnline,
	}); err != nil {
		b.Logger.Errorf("Failed to set presence: %s", err)
	}
}
