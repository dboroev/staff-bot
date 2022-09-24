package config

import "github.com/spf13/viper"

type Config struct {
	TelegramToken string
	DbUrl         string `mapstructure:"db_url"`

	Messages Messages
}

type Messages struct {
	Errors
	Responses
}

type Errors struct {
	Default string `mapstructure:"default"`
}

type Responses struct {
	Start string `mapstructure:"start"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("token")

	return nil
}
