package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/bot-template"
	"github.com/disgoorg/bot-template/commands"
	"github.com/disgoorg/bot-template/components"
	"github.com/disgoorg/bot-template/handlers"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/log"
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
	h.HandleCommand("/test", commands.TestHandler)
	h.HandleAutocomplete("/test", commands.TestAutocompleteHandler)
	h.HandleCommand("/version", commands.VersionHandler(b))
	h.HandleComponent("test_button", components.TestComponent)

	b.SetupBot(h, bot.NewListenerFunc(b.OnReady), handlers.MessageHandler(b))

	if *shouldSyncCommands {
		if cfg.DevMode {
			logger.Warn("Syncing commands in dev mode")
			_, err = b.Client.Rest().SetGuildCommands(b.Client.ApplicationID(), cfg.DevGuildID, commands.Commands)
		} else {
			logger.Info("Syncing commands")
			_, err = b.Client.Rest().SetGlobalCommands(b.Client.ApplicationID(), commands.Commands)
		}
		if err != nil {
			logger.Errorf("failed to sync commands: %s", err.Error())
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	if err = b.Client.OpenGateway(ctx); err != nil {
		b.Logger.Errorf("Failed to connect to gateway: %s", err)
	}
	defer func() {
		cctx, ccancel := context.WithTimeout(context.Background(), 10)
		defer ccancel()
		b.Client.Close(cctx)
	}()

	b.Logger.Info("Bot is running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
	b.Logger.Info("Shutting down...")
}
