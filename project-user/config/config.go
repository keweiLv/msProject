package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var C = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

func InitConfig() *Config {
	conf := &Config{
		viper: viper.New(),
	}
	workDur, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDur + "/config")
	conf.viper.AddConfigPath("/etc/msProject/user")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf.ReadServerConfig()
	return conf
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}
