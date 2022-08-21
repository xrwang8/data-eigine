package conf

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
	Log      LogConfig      `yaml:"log"`
}

type AppConfig struct {
	Port        string `yaml:"port"`
	Environment string `mapstructure:"env"`
}

type LogConfig struct {
	LogOutput string `mapstructure:"log_output"`
	LOG_LEVEL string `mapstructure:"log_level"`
}
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

func NewConf() Config {
	conf := Config{}
	viper.SetConfigFile("app.yml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("can't found config file %s", err)
		} else {
			log.Fatalf("error parsing configuration file %s", err)
		}
	}
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatal(err)
	}
	return conf
}
