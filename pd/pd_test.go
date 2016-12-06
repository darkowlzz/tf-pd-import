package pd

import (
	"errors"
	"reflect"
	"testing"

	"github.com/PagerDuty/go-pagerduty"
)

type FakePDClient struct {
	EscalationPolicyCount int
	ServiceCount          int
	Err                   error
}

func (f FakePDClient) GetEscalationPolicies() ([]pagerduty.EscalationPolicy, error) {
	var escalationPolicies []pagerduty.EscalationPolicy
	for i := 0; i < f.EscalationPolicyCount; i++ {
		apiobj := pagerduty.APIObject{ID: string(i)}
		newEP := pagerduty.EscalationPolicy{APIObject: apiobj}
		escalationPolicies = append(escalationPolicies, newEP)
	}
	return escalationPolicies, f.Err
}

func (f FakePDClient) GetServices() ([]pagerduty.Service, error) {
	var services []pagerduty.Service
	for i := 0; i < f.ServiceCount; i++ {
		apiobj := pagerduty.APIObject{ID: string(i)}
		newService := pagerduty.Service{APIObject: apiobj}
		services = append(services, newService)
	}
	return services, f.Err
}

func TestGetEscalationPolicies(t *testing.T) {
	cases := []struct {
		f                          FakePDClient
		expectedEscalationPolicies int
		expectedErr                error
	}{
		{
			f: FakePDClient{
				EscalationPolicyCount: 5,
				Err: nil,
			},
			expectedEscalationPolicies: 5,
			expectedErr:                nil,
		},
		{
			f: FakePDClient{
				EscalationPolicyCount: 12,
				Err: nil,
			},
			expectedEscalationPolicies: 12,
			expectedErr:                nil,
		},
		{
			f: FakePDClient{
				EscalationPolicyCount: 0,
				Err: errors.New("TCP timeout"),
			},
			expectedEscalationPolicies: 0,
			expectedErr:                errors.New("TCP timeout"),
		},
	}

	for _, c := range cases {
		eps, err := c.f.GetEscalationPolicies()
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Fatalf("Expected err to be %s but got %s", c.expectedErr, err)
		}

		if len(eps) != c.expectedEscalationPolicies {
			t.Fatalf("Expected %d but got %d escalation policies", c.expectedEscalationPolicies, len(eps))
		}
	}
}

func TestGetServices(t *testing.T) {
	cases := []struct {
		f                FakePDClient
		expectedServices int
		expectedErr      error
	}{
		{
			f: FakePDClient{
				ServiceCount: 10,
				Err:          nil,
			},
			expectedServices: 10,
			expectedErr:      nil,
		},
		{
			f: FakePDClient{
				ServiceCount: 3,
				Err:          errors.New("TCP timeout"),
			},
			expectedServices: 3,
			expectedErr:      errors.New("TCP timeout"),
		},
	}

	for _, c := range cases {
		services, err := c.f.GetServices()
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Fatalf("Expected err to be %s but got %s", c.expectedErr, err)
		}

		if len(services) != c.expectedServices {
			t.Fatalf("Expected %d but got %d services", c.expectedServices, len(services))
		}
	}
}
