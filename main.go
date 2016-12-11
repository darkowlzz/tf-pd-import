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
	authtoken := tf.GetStringConfig("authtoken")
	tfclient := tf.NewTf(authtoken)

	// Check if terraformBin is defined and use the binary if defined
	tfBin := tf.GetStringConfig("terraformBin")
	if tfBin != "" {
		tfclient.TerraformBin = tfBin
	}

	tf.ImportServices(tfclient)
	tf.ImportEscalationPolicies(tfclient)
}
