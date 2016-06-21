package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/chyeh/viper"
)

func getFileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func loadConfigFile() {
	cfgPath := viper.GetString("config")
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalln("Configuration file [", cfgPath, "] doesn't exist")
	}

	viper.SetConfigName(getFileNameWithoutExtension(cfgPath))

	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Error:", err)
	}

}
