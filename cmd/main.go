package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/log"

	"github.com/disgoorg/bot-template"
	"github.com/disgoorg/bot-template/commands"
	"github.com/disgoorg/bot-template/components"
	"github.com/disgoorg/bot-template/handlers"
)

var (
	shouldSyncCommands *bool
	version            = "dev"
)

func init() {
	shouldSyncCommands = flag.Bool("sync-commands", false, "Whether to sync commands to discord")
	flag.Parse()
}

func main() {
	cfg, err := dbot.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	logger := log.New(log.Ldate | log.Ltime | log.Lshortfile)
	logger.SetLevel(cfg.LogLevel)
	logger.Infof("Starting bot version: %s", version)
	logger.Infof("Syncing commands? %t", *shouldSyncCommands)

	b := dbot.New(logger, version, *cfg)

	h := handler.New()
	h.Command("/test", commands.TestHandler)
	h.Autocomplete("/test", commands.TestAutocompleteHandler)
	h.Command("/version", commands.VersionHandler(b))
	h.Component("test_button", components.TestComponent)

	b.SetupBot(h, bot.NewListenerFunc(b.OnReady), handlers.MessageHandler(b))

	if *shouldSyncCommands {
		if cfg.DevMode {
			logger.Info("Syncing commands in dev mode")
			_, err = b.Client.Rest().SetGuildCommands(b.Client.ApplicationID(), cfg.DevGuildID, commands.Commands)
		} else {
			logger.Info("Syncing commands")
			_, err = b.Client.Rest().SetGlobalCommands(b.Client.ApplicationID(), commands.Commands)
		}
		if err != nil {
			logger.Errorf("failed to sync commands: %s", err.Error())
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = b.Client.OpenGateway(ctx); err != nil {
		b.Logger.Errorf("Failed to connect to gateway: %s", err)
	}
	defer b.Client.Close(context.TODO())

	b.Logger.Info("Bot is running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
	b.Logger.Info("Shutting down...")
}
