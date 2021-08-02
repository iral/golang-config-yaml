package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	port := viper.Get("server.port")
	username := viper.Get("database.username")
	password := viper.Get("database.password")

	fmt.Println(port)
	fmt.Println(username)
	fmt.Println(password)

	// Viper supports the ability to have your application live read a config file while running.
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
