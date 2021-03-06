package tf

import (
	"errors"
	"reflect"
	"testing"
)

func TestTfClient(t *testing.T) {
	var called bool
	var returnError error

	// A mock function that would be called by ImportEscalationPolicy and
	// ImportService
	fakeImportRes := func(token, tfBin, resType, name, id string) error {
		called = !called
		return returnError
	}

	tf := TfClient{pdToken: "AAAA", importRes: fakeImportRes,
		TerraformBin: "terraform"}

	cases := []struct {
		f            func(id, name string) error
		expectedCall bool
		expectedErr  error
	}{
		{
			f:            tf.ImportEscalationPolicy,
			expectedCall: true,
			expectedErr:  nil,
		},
		{
			f:            tf.ImportService,
			expectedCall: true,
			expectedErr:  nil,
		},
		{
			f:            tf.ImportEscalationPolicy,
			expectedCall: true,
			expectedErr:  errors.New("TCP timeout"),
		},
	}

	for _, c := range cases {
		// Reset values in every iteration
		called = false
		returnError = c.expectedErr
		err := c.f("BBBB", "YYY")
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Fatalf("Expected err to be %s but got %s", c.expectedErr, err)
		}
		if called != c.expectedCall {
			t.Fatalf("Expected called to be %d but got %d", c.expectedCall, called)
		}
	}
}

func TestGetResourceName(t *testing.T) {
	cases := []struct {
		resType       string
		name          string
		expectedValue string
		expectedErr   error
	}{
		{
			resType:       EscalationPolicyPrefix,
			name:          "MMMM",
			expectedValue: EscalationPolicyPrefix + "." + "MMMM",
			expectedErr:   nil,
		},
		{
			resType:       ServicePrefix,
			name:          "KKKK",
			expectedValue: ServicePrefix + "." + "KKKK",
			expectedErr:   nil,
		},
		{
			resType:       "Foo",
			name:          "LLLL",
			expectedValue: "",
			expectedErr:   errors.New("Unknown resource type Foo"),
		},
	}

	for _, c := range cases {
		val, err := getResourceName(c.resType, c.name)
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Fatalf("Expected err to be %s but got %s", c.expectedErr, err)
		}
		if val != c.expectedValue {
			t.Fatalf("Expected resourceName to be %s but got %s", c.expectedValue, val)
		}
	}
}

func TestTerraformImport(t *testing.T) {
	// TODO: Find ways to test this
}

func TestImportResource(t *testing.T) {
	var getResourceNameCalled bool
	var terraformImportCalled bool
	var getResourceNameErr error
	var terraformImportErr error

	getResourceName = func(resType, name string) (string, error) {
		getResourceNameCalled = !getResourceNameCalled
		return "", getResourceNameErr
	}

	terraformImport = func(token, tfBin, resName, id string) error {
		terraformImportCalled = !terraformImportCalled
		return terraformImportErr
	}

	cases := []struct {
		getResErr              error
		tfImportErr            error
		expectedGetResCalled   bool
		expectedTfImportCalled bool
		expectedErr            error
	}{
		{
			getResErr:              nil,
			tfImportErr:            nil,
			expectedGetResCalled:   true,
			expectedTfImportCalled: true,
			expectedErr:            nil,
		},
		{
			getResErr:              errors.New("TCP error"),
			tfImportErr:            nil,
			expectedGetResCalled:   true,
			expectedTfImportCalled: false,
			expectedErr:            errors.New("TCP error"),
		},
		{
			getResErr:              nil,
			tfImportErr:            errors.New("TCP error"),
			expectedGetResCalled:   true,
			expectedTfImportCalled: true,
			expectedErr:            errors.New("TCP error"),
		},
	}

	for _, c := range cases {
		getResourceNameCalled = false
		terraformImportCalled = false
		getResourceNameErr = c.getResErr
		terraformImportErr = c.tfImportErr
		err := importResource("token1", "tf", "res", "name", "id")
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Fatalf("Expected err to be %s but got %s", c.expectedErr, err)
		}
		if getResourceNameCalled != c.expectedGetResCalled {
			t.Fatalf(
				"Expected getResourceNameCalled to be %d but got %d",
				c.expectedGetResCalled, getResourceNameCalled,
			)
		}
		if terraformImportCalled != c.expectedTfImportCalled {
			t.Fatalf(
				"Expected terraformImportCalled to be %d but got %d",
				c.expectedTfImportCalled, terraformImportCalled,
			)
		}
	}
}
