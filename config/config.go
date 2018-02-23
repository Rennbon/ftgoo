package config

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	TaskCenter  Taskcenter
	Certificate Certificate
}

var filePath string

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
}
func LoadConfig() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}
	cfg := &Config{}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	return cfg, nil
}

func CheckConfig(c *Config, cnames []string) error {
	s := reflect.ValueOf(&c).Elem()
	for _, v := range cnames {
		if s.FieldByName(v).Interface() == nil {
			return fmt.Errorf("%v is not find in config", v)
		}
	}
	return nil
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

type Certificate struct {
	CertFile string
	KeyFile  string
}
