package main

import (
	"github.com/darkowlzz/tf-pd-import/tf"
	"github.com/spf13/viper"
)

func configInit() {
	// initialize viper and read config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	configInit()
	// type assert config to be string
	authtoken := tf.GetConfig("authtoken").(string)
	tfclient := tf.NewTf(authtoken)

	tf.ImportServices(tfclient)
	tf.ImportEscalationPolicies(tfclient)
}
