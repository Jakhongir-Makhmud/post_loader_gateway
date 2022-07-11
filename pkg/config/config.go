package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
}

type cfg struct {
	c *viper.Viper
}

func NewConfig() Config {
	v := viper.New()

	v.SetConfigName("apiconfig")
	v.SetConfigType("json")

	v.AddConfigPath("./config")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.WatchConfig()

	return &cfg{c: v}
}

func (c *cfg) GetString(key string) string {
	return c.c.GetString(key)
}

func (c *cfg) GetInt(key string) int {
	return c.c.GetInt(key)
}