package main

import "github.com/spf13/viper"

func init() {
	viper.GetViper().SetConfigFile()
}
