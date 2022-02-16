package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/DisgoOrg/bot-template/internal/types"
	"github.com/DisgoOrg/bot-template/modules"
	"github.com/DisgoOrg/log"
)

var (
	shouldSyncCommands *bool
	shouldSyncDBTables *bool
	exitAfterSync      *bool
	version            = "dev"
)

func init() {
	shouldSyncCommands = flag.Bool("sync-modules", false, "Whether to sync commands to discord")
	shouldSyncDBTables = flag.Bool("sync-db", false, "Whether to sync the database tables")
	exitAfterSync = flag.Bool("exit-after", false, "Whether to exit after syncing commands and database tables")
	flag.Parse()
}

func main() {
	var err error
	logger := log.New(log.Ldate | log.Ltime | log.Lshortfile)
	bot := types.Bot{
		Logger: logger,
	}
	bot.Logger.Infof("Starting bot version: %s", version)
	bot.Logger.Infof("Syncing commands? %v", *shouldSyncCommands)
	bot.Logger.Infof("Syncing DB tables? %v", *shouldSyncDBTables)
	bot.Logger.Infof("Exiting after syncing? %v", *exitAfterSync)
	defer bot.Logger.Info("Shutting down bot...")

	if err = bot.LoadConfig(); err != nil {
		bot.Logger.Fatal("Failed to load config: ", err)
	}
	logger.SetLevel(bot.Config.LogLevel)

	bot.LoadModules(modules.Modules)

	if err = bot.SetupBot(); err != nil {
		bot.Logger.Fatal("Failed to setup bot: ", err)
	}
	defer bot.Bot.Close(context.TODO())

	if *shouldSyncCommands {
		if err = bot.SyncCommands(); err != nil {
			bot.Logger.Fatal("Failed to sync modules: ", err)
		}
	}

	if err = bot.SetupDatabase(*shouldSyncDBTables); err != nil {
		bot.Logger.Fatal("Failed to setup database: ", err)
	}
	defer bot.DB.Close()

	if *exitAfterSync {
		bot.Logger.Infof("Exiting after syncing commands and database tables")
		os.Exit(0)
	}

	if err = bot.StartBot(); err != nil {
		bot.Logger.Fatal("Failed to start bot: ", err)
	}

	bot.Logger.Info("Bot is running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}
