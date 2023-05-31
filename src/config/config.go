package config

import (
	"apixyz/src/apixyz/util"
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
func Realm() string {
	return viper.GetString("realm")
}
func SuperSecretPassword() string {
	return util.DecryptAES(viper.GetString("token.supersecretpassword"))
}

// returns Timeout duration value
func Timeout() int {
	return viper.GetInt("timeout")
}
