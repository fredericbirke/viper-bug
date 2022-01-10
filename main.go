package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetEnvPrefix("TEST_PREFIX")
	if err := viper.ReadInConfig(); err != nil {
		var configNotFoundErr viper.ConfigFileNotFoundError
		if errors.Is(err, configNotFoundErr) {
			panic("cannot read config: " + err.Error())
		}
	}

	foundSettings := viper.AllSettings()
	fmt.Println("Found Settings:")
	for s, _ := range foundSettings {
		fmt.Println(s)
	}

	// This should be 'bar'
	value := viper.GetString("foo")
	fmt.Println("Requested value:")
	fmt.Println(value)
}
