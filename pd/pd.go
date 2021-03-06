package pd

import (
	"github.com/PagerDuty/go-pagerduty"
)

type PDClient interface {
	GetEscalationPolicies() ([]pagerduty.EscalationPolicy, error)
	GetServices() ([]pagerduty.Service, error)
}

type TfPDClient struct {
	client pagerduty.Client
}

// GetEscalationPolicies fetches and returns all the Pagerduty Escalation Policies.
func (c TfPDClient) GetEscalationPolicies() ([]pagerduty.EscalationPolicy, error) {
	var opts pagerduty.ListEscalationPoliciesOptions
	var escalationPolicies []pagerduty.EscalationPolicy
	for {
		opts.APIListObject.Offset = uint(len(escalationPolicies))
		if policyResponse, err := c.client.ListEscalationPolicies(opts); err != nil {
			return escalationPolicies, err
		} else {
			for _, policy := range policyResponse.EscalationPolicies {
				escalationPolicies = append(escalationPolicies, policy)
			}
			if !policyResponse.APIListObject.More {
				break
			}
		}
	}
	return escalationPolicies, nil
}

// GetServices fetches and returns all the Pagerduty Services.
func (c TfPDClient) GetServices() ([]pagerduty.Service, error) {
	var opts pagerduty.ListServiceOptions
	var services []pagerduty.Service
	for {
		opts.APIListObject.Offset = uint(len(services))
		if serviceResponse, err := c.client.ListServices(opts); err != nil {
			return services, err
		} else {
			for _, service := range serviceResponse.Services {
				services = append(services, service)
			}
			if !serviceResponse.APIListObject.More {
				break
			}
		}
	}
	return services, nil
}
