package bot_template

import (
	"errors"
	"os"

	"github.com/disgoorg/disgo/json"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if os.IsNotExist(err) {
		if file, err = os.Create("config.json"); err != nil {
			return nil, err
		}
		var data []byte
		if data, err = json.MarshalIndent(Config{}, "", "\\t"); err != nil {
			return nil, err
		}
		if _, err = file.Write(data); err != nil {
			return nil, err
		}
		return nil, errors.New("config.json not found, created new one")
	} else if err != nil {
		return nil, err
	}

	var cfg Config
	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

type Config struct {
	DevMode    bool         `json:"dev_mode"`
	DevGuildID snowflake.ID `json:"dev_guild_id"`
	LogLevel   log.Level    `json:"log_level"`
	Token      string       `json:"token"`
}
