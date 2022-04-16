package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConf
	Database DatabaseConf
}

type ServerConf struct {
	Port int
}

type DatabaseConf struct {
	Host       string
	Port       int
	DBName     string
	DBUser     string
	DBPassword string
	Timezone   string
}

func LoadConfig(env string) (*Config, error) {
	viper.SetConfigName(fmt.Sprintf("%s-config", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/auth-micro/")
	viper.AddConfigPath("$HOME/.auth-micro/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("../../config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	conf := new(Config)
	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}
	overrideUsingEnvVars(conf)
	return conf, nil
}

// Added workaround due to issues with environment variables in Viper
// https://github.com/spf13/viper/issues/761
func overrideUsingEnvVars(config *Config) {
	if host, present := os.LookupEnv("DB_HOST"); present {
		config.Database.Host = host
	}
}
