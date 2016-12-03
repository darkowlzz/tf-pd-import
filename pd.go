package main

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/spf13/viper"
)

type PDClient interface {
	GetEscalationPolicies(string) ([]pagerduty.EscalationPolicy, error)
}

type TfPDClient struct{}

// Returns all the Pagerduty Escalation Policies
func (c TfPDClient) GetEscalationPolicies(token string) ([]pagerduty.EscalationPolicy, error) {
	var opts pagerduty.ListEscalationPoliciesOptions
	client := pagerduty.NewClient(token)
	var escalationPolicies []pagerduty.EscalationPolicy
	for {
		opts.APIListObject.Offset = uint(len(escalationPolicies))
		if eps, err := client.ListEscalationPolicies(opts); err != nil {
			panic(err)
		} else {
			for _, p := range eps.EscalationPolicies {
				escalationPolicies = append(escalationPolicies, p)
			}
			if ! eps.APIListObject.More {
				break
			}
		}
	}
	return escalationPolicies, nil
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	authtoken := viper.GetString("authtoken")
	tfclient := TfPDClient{}
	tfclient.GetEscalationPolicies(authtoken)
}
