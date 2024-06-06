package conf

import (
	"errors"
	"log/slog"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	App AppConfig `mapstructure:"app"`
}

type AppConfig struct {
	Port      string `mapstructure:"port"`
	MongoUri  string `mapstructure:"mongoUri"`
	DBName    string `mapstructure:"dbname"`
	UserColl  string `mapstructure:"userColl"`
	HotelColl string `mapstructure:"hotelColl"`
	RoomColl  string `mapstructure:"roomColl"`
	Secert    string `mapstructure:"secert"`
	JWTSecert string `mapstructure:"jwtSecert"`
}

func ParseConfig() error {
	viper.SetConfigName("hotel")
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
