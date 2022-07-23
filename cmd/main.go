package main

import (
	"flag"

	"github.com/disgoorg/bot-template"
	"github.com/disgoorg/bot-template/commands"
	"github.com/disgoorg/bot-template/components"
	"github.com/disgoorg/bot-template/handlers"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
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
	cfg, err := bot_template.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	logger := log.New(log.Ldate | log.Ltime | log.Lshortfile)
	logger.SetLevel(cfg.LogLevel)
	logger.Infof("Starting bot version: %s", version)
	logger.Infof("Syncing commands? %v", *shouldSyncCommands)

	b := bot_template.New(logger, version, *cfg)
	b.SetupBot(handlers.MessageHandler(b))
	b.Handler.AddCommands(commands.TestCommand(b))
	b.Handler.AddComponents(components.TestComponent(b))

	if *shouldSyncCommands {
		var guildIDs []snowflake.ID
		if cfg.DevMode {
			guildIDs = append(guildIDs, cfg.DevGuildID)
		}
		b.Handler.SyncCommands(b.Client, guildIDs...)
	}

	b.StartAndBlock()
}
