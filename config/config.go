package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func New(configFile, dbYaml string) {
	viper.SetConfigFile(configFile)
	// viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("read config file %s failed: %v\n", viper.ConfigFileUsed(), err)
		os.Exit(1)
	}
	fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())

	viper.SetConfigFile(dbYaml)
	// viper.AddConfigPath(".")
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("merge config file[%s], failed[%s]\n", dbYaml, err)

		os.Exit(1)
	}

}
