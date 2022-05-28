package main

import (
	"flag"

	"github.com/disgoorg/bot-template/commands"
	"github.com/disgoorg/bot-template/components"
	"github.com/disgoorg/bot-template/handlers"
	"github.com/disgoorg/bot-template/tbot"
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
	cfg, err := tbot.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	logger := log.New(log.Ldate | log.Ltime | log.Lshortfile)
	logger.SetLevel(cfg.LogLevel)
	logger.Infof("Starting bot version: %s", version)
	logger.Infof("Syncing commands? %v", *shouldSyncCommands)

	b := tbot.New(logger, version, *cfg)
	b.SetupBot(handlers.MessageHandler(b))
	b.SetupCommands(*shouldSyncCommands, commands.TestCommand)
	b.SetupComponents(components.TestComponent)
	b.StartAndBlock()
}
