package config

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	TaskCenter Taskcenter
}

var filePath string

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	//viper.AddConfigPath("/Users/zhubin/go/src/ftgoo/config/")
}
func LoadConfig() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {

		log.Fatalf("Error reading config file, %s", err)
	}
	cfg := &Config{}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

func WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println("Config file changed:", e.Op)
		fmt.Println("Config file changed:", e.String())
	})
}

type Taskcenter struct {
	Addr      string
	Timeout   time.Duration
	PoolLimit int
	Database  string
}
