package tf

import (
	"testing"

	"github.com/spf13/viper"
)

type FakeTerraformClient struct {
	importServiceCallCount          int
	importEscalationPolicyCallCount int
	services                        []map[string]string
	policies                        []map[string]string
}

func (f *FakeTerraformClient) ImportService(id, name string) error {
	// importServiceCallCount is used to measure the number of times
	// ImportService is called.
	f.importServiceCallCount++
	// services is used to store all the services and ids that ImportService
	// received for importing.
	f.services = append(f.services, map[string]string{"name": name, "id": id})
	return nil
}

func (f *FakeTerraformClient) ImportEscalationPolicy(id, name string) error {
	// importEscalationPolicyCallCount is used to measure the number of times
	// ImportEscalationPolicy is called.
	f.importEscalationPolicyCallCount++
	// policies is used to store all the policies and ids that
	// ImportEscalationPolicy received for importing.
	f.policies = append(f.policies, map[string]string{"name": name, "id": id})
	return nil
}

func TestImportServices(t *testing.T) {
	cases := []struct {
		tf                *FakeTerraformClient
		mockServices      []map[string]string
		expectedCallCount int
	}{
		{
			tf: &FakeTerraformClient{},
			mockServices: []map[string]string{
				{
					"name": "foo0", "id": "id0",
				},
				{
					"name": "foo1", "id": "id1",
				},
			},
			expectedCallCount: 2,
		},
		{
			tf: &FakeTerraformClient{},
			mockServices: []map[string]string{
				{
					"name": "qqq1", "id": "qid1",
				},
				{
					"name": "qqq2", "id": "qid2",
				},
				{
					"name": "qqq3", "id": "qid3",
				},
			},
			expectedCallCount: 3,
		},
		{
			tf:                &FakeTerraformClient{},
			mockServices:      nil,
			expectedCallCount: 0,
		},
	}

	for _, c := range cases {
		// Create mock config for test case
		s := make([]interface{}, len(c.mockServices))
		for i := 0; i < len(c.mockServices); i++ {
			s[i] = map[interface{}]interface{}{
				"name": c.mockServices[i]["name"],
				"id":   c.mockServices[i]["id"],
			}
		}
		// Set the mock config in viper
		viper.Set("services", s)

		ImportServices(c.tf)

		// Check if ImportService was called expected number of times
		if c.tf.importServiceCallCount != c.expectedCallCount {
			t.Fatalf("Expected ImportService to be called %d times, but got %d",
				c.expectedCallCount, c.tf.importServiceCallCount)
		}

		// Check if the expected services and ids were passed to ImportService
		for _, expectedService := range c.tf.services {
			serviceFound := false
			for _, service := range c.mockServices {
				if expectedService["name"] == service["name"] &&
					expectedService["id"] == service["id"] {
					serviceFound = true
					break
				}
			}
			if !serviceFound {
				t.Fatalf("Expected %v to be passed to ImportService",
					expectedService)
			}
		}
	}
}

func TestImportEscalationPolicies(t *testing.T) {
	cases := []struct {
		tf                *FakeTerraformClient
		mockPolicies      []map[string]string
		expectedCallCount int
	}{
		{
			tf: &FakeTerraformClient{},
			mockPolicies: []map[string]string{
				{
					"name": "foo0", "id": "id0",
				},
				{
					"name": "foo1", "id": "id1",
				},
			},
			expectedCallCount: 2,
		},
		{
			tf: &FakeTerraformClient{},
			mockPolicies: []map[string]string{
				{
					"name": "qqq1", "id": "qid1",
				},
				{
					"name": "qqq2", "id": "qid2",
				},
				{
					"name": "qqq3", "id": "qid3",
				},
			},
			expectedCallCount: 3,
		},
		{
			tf:                &FakeTerraformClient{},
			mockPolicies:      nil,
			expectedCallCount: 0,
		},
	}

	for _, c := range cases {
		// Create mock config for test case
		s := make([]interface{}, len(c.mockPolicies))
		for i := 0; i < len(c.mockPolicies); i++ {
			s[i] = map[interface{}]interface{}{
				"name": c.mockPolicies[i]["name"],
				"id":   c.mockPolicies[i]["id"],
			}
		}

		// Set the mock config in viper
		viper.Set("policies", s)

		ImportEscalationPolicies(c.tf)

		// Check if ImportEscalationPolicy was called expected number of times
		if c.tf.importEscalationPolicyCallCount != c.expectedCallCount {
			t.Fatalf(
				"Expected EscalationPolicy to be called %d times, but got %d",
				c.expectedCallCount, c.tf.importEscalationPolicyCallCount,
			)
		}

		// Check if the expected policy and ids were passed to
		// ImportEscalationPolicy
		for _, expectedPolicy := range c.tf.policies {
			policyFound := false
			for _, policy := range c.mockPolicies {
				if expectedPolicy["name"] == policy["name"] &&
					expectedPolicy["id"] == policy["id"] {
					policyFound = true
					break
				}
			}
			if !policyFound {
				t.Fatalf("Expected %v to be passed to ImportEscalationPolicy",
					expectedPolicy)
			}
		}
	}
}
