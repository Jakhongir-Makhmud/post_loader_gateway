package config

import (
	"os"
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
	v.AddConfigPath(getConfigPath())
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

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

func getConfigPath() (path string) {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	slice := strings.Split(wd, "api-gateway-iman")
	path = slice[0] + "api-gateway-iman/config"
	return
}
