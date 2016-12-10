package tf

import (
	"github.com/spf13/viper"
)

// GetConfig reads a given property from viper config and returns an interface
var GetConfig = func(prop string) interface{} {
	return viper.Get(prop)
}

// GetStringConfig reads given property from viper config and returns a string
// value.
var GetStringConfig = func(prop string) string {
	return viper.GetString(prop)
}

// ImportServices imports all the Pagerduty Services in config file
func ImportServices(tclient TerraformClient) error {
	services := GetConfig("services")
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

// ImportEscalationPolicies imports all the Pagerduty EscalationPolicies in
// config file
func ImportEscalationPolicies(tclient TerraformClient) error {
	policies := GetConfig("policies")
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
