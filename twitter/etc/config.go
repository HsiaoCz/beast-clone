package etc

import (
	"errors"
	"log/slog"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// why we use config but not use env
// because the env need under the floder

var Conf = new(Config)

type Config struct {
	App AppConf `mapstructure:"app"`
}

type AppConf struct {
	Port        string `mapstructure:"port"`
	MongoUri    string `mapstructure:"mongoUri"`
	DBName      string `mapstructure:"dbname"`
	UserColl    string `mapstructure:"userColl"`
	Secert      string `mapstructure:"secret"`
	JWTSecert   string `mapstructure:"jwtSecret"`
	PostColl    string `mapstructure:"postColl"`
	CommentColl string `mapstructure:"commentColl"`
}

func ParseConfig() error {
	viper.SetConfigName("twitter")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return errors.New("viper read in config error")
	}

	if err := viper.Unmarshal(Conf); err != nil {
		return errors.New("viper unmarshal error")
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		slog.Info("config changed...")
		if err := viper.Unmarshal(Conf); err != nil {
			slog.Error("viper reunmarshal error", "error message", err)
			return
		}
	})
	return nil
}
