package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Environment() string {
	env := viper.GetString("environment")
	fmt.Println(env)

	return env
}

func Hostname() string {
	host := viper.GetString("host")
	fmt.Println(host)

	return host
}
func DetermineListenAddress() string {
	port := viper.GetString("port")
	return ":" + port
}

// returns Timeout duration value
func Timeout() int {
	return viper.GetInt("timeout")
}
