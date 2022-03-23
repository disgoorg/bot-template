package types

import (
	"os"

	"github.com/disgoorg/disgo/json"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake"
	"github.com/pkg/errors"
)

func (b *Bot) LoadConfig() error {
	b.Logger.Info("Loading config...")
	file, err := os.Open("config.json")
	if os.IsNotExist(err) {
		if file, err = os.Create("config.json"); err != nil {
			return err
		}
		var data []byte
		if data, err = json.Marshal(Config{}); err != nil {
			return err
		}
		if _, err = file.Write(data); err != nil {
			return err
		}
		return errors.New("config.json not found, created new one")
	} else if err != nil {
		return err
	}

	var cfg Config
	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		return err
	}
	b.Config = cfg
	return nil
}

type Config struct {
	DevMode    bool                `json:"dev_mode"`
	DevGuildID snowflake.Snowflake `json:"dev_guild_id"`
	LogLevel   log.Level           `json:"log_level"`
	Token      string              `json:"token"`

	Database Database `json:"database"`
}

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}
