package main

import (
	"testing"

	"github.com/PagerDuty/go-pagerduty"
)

type FakePDClient struct {
	EscalationPolicyCount int
}

func (f FakePDClient) GetEscalationPolicies(token string) ([]pagerduty.EscalationPolicy, error) {
	var escalationPolicies []pagerduty.EscalationPolicy
	for i := 0; i < f.EscalationPolicyCount; i++ {
		apiobj := pagerduty.APIObject{ID: string(i)}
		newEP := pagerduty.EscalationPolicy{APIObject: apiobj}
		escalationPolicies = append(escalationPolicies, newEP)
	}
	return escalationPolicies, nil
}

func TestGetEscalationPolicies(t *testing.T) {
	cases := []struct {
		f                          FakePDClient
		expectedEscalationPolicies int
		expectedErr                error
	}{
		{
			f: FakePDClient{EscalationPolicyCount: 5},
			expectedEscalationPolicies: 5,
			expectedErr:                nil,
		},
		{
			f: FakePDClient{EscalationPolicyCount: 12},
			expectedEscalationPolicies: 12,
			expectedErr:                nil,
		},
	}

	for _, c := range cases {
		eps, err := c.f.GetEscalationPolicies("dummytoken")
		if err != c.expectedErr {
			t.Fatalf("Expected err to be nil but it was %s", err)
		}

		if len(eps) != c.expectedEscalationPolicies {
			t.Fatalf("Expected %d but got %d escalation policies", c.expectedEscalationPolicies, len(eps))
		}
	}
}
