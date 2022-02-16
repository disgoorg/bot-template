package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/DisgoOrg/log"
	"github.com/your-name/your-repo/commands"
	"github.com/your-name/your-repo/internal/types"
)

var (
	shouldSyncCommands *bool
	shouldSyncDBTables *bool
)

func init() {
	shouldSyncCommands = flag.Bool("sync-commands", false, "Whether to sync commands")
	shouldSyncDBTables = flag.Bool("sync-db", false, "Whether to sync the database tables")
	flag.Parse()
}

func main() {
	var err error
	logger := log.New(log.Ldate | log.Ltime | log.Lshortfile)
	bot := types.Bot{
		Logger: logger,
	}

	if err = bot.LoadConfig(); err != nil {
		bot.Logger.Fatal("Failed to load config: ", err)
	}
	logger.SetLevel(bot.Config.LogLevel)

	bot.LoadCommands(commands.Commands)

	if err = bot.SetupBot(); err != nil {
		bot.Logger.Fatal("Failed to setup bot: ", err)
	}

	if *shouldSyncCommands {
		if err = bot.SyncCommands(); err != nil {
			bot.Logger.Fatal("Failed to sync commands: ", err)
		}
	}

	if err = bot.SetupDatabase(*shouldSyncDBTables); err != nil {
		bot.Logger.Fatal("Failed to setup database: ", err)
	}

	if err = bot.StartBot(); err != nil {
		bot.Logger.Fatal("Failed to start bot: ", err)
	}

	bot.Logger.Info("Bot is running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-s
	bot.Logger.Info("Shutting down bot...")
	bot.Bot.Close(context.TODO())
	_ = bot.DB.Close()
}
