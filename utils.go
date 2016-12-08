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

// getConfig reads a given property from config and returns an interface
var getConfig = func(prop string) interface{} {
	return viper.Get(prop)
}

// importServices imports all the Pagerduty Services in config file
func importServices(tclient tf.TerraformClient) error {
	services := getConfig("services")
	serviceList := services.([]interface{})
	for _, val := range serviceList {
		service := val.(map[interface{}]interface{})
		err := tclient.ImportService(service["id"].(string), service["name"].(string))
		if err != nil {
			return err
		}
	}
	return nil
}

// importEscalationPolicies imports all the Pagerduty EscalationPolicies in
// config file
func importEscalationPolicies(tclient tf.TerraformClient) error {
	policies := getConfig("policies")
	policyList := policies.([]interface{})
	for _, val := range policyList {
		policy := val.(map[interface{}]interface{})
		err := tclient.ImportEscalationPolicy(policy["id"].(string), policy["name"].(string))
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
    configInit()
	// type assert config to be string
	authtoken := getConfig("authtoken").(string)
	tfclient := tf.NewTf(authtoken)

	importServices(tfclient)
	importEscalationPolicies(tfclient)
}
