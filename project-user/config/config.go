package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/keweiLv/project-common/logs"
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
	conf.InitZapLog()
	return conf
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}

func (c *Config) InitZapLog() {
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("zap.maxsize"),
		MaxAge:        c.viper.GetInt("zap.max_age"),
		MaxBackups:    c.viper.GetInt("zap.max_backups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		panic(err)
	}
}

func (c *Config) ReadRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     c.viper.GetString("redis.host") + ":" + c.viper.GetString("redis.port"),
		Password: c.viper.GetString("redis.password"),
		DB:       c.viper.GetInt("redis.db"),
	}
}
