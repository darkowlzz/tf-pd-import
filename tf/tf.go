package tf

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

const (
	EscalationPolicyPrefix = "pagerduty_escalation_policy"
	ServicePrefix          = "pagerduty_service"
)

type TerraformClient interface {
	ImportService(id, name string) error
	ImportEscalationPolicy(id, name string) error
}

type TfClient struct {
	pdToken   string
	importRes func(token, resType, name, id string) error
}

// ImportEscalationPolicy imports Pagerduty Escalation Policy as a terraform
// resource, given a resource id and name.
func (c TfClient) ImportEscalationPolicy(id, name string) error {
	return c.importRes(c.pdToken, EscalationPolicyPrefix, name, id)
}

// ImportService imports Pagerduty Service as a terraform resource, given a
// resource id and name.
func (c TfClient) ImportService(id, name string) error {
	return c.importRes(c.pdToken, ServicePrefix, name, id)
}

// getResourceName combines a resource type with a given name to return a string
// that could be used with terraform import command.
func getResourceName(resType, name string) (string, error) {
	switch resType {
	case EscalationPolicyPrefix:
		return strings.Join([]string{EscalationPolicyPrefix, name}, "."), nil
	case ServicePrefix:
		return strings.Join([]string{ServicePrefix, name}, "."), nil
	default:
		return "", errors.New("Unknown resource type " + resType)
	}
}

// terraformImport uses terraform import command to import an existing resource,
// provided a provider's token, resourceName and resource terraform ID.
func terraformImport(token, resourceName, id string) error {
	var err error
	cmd1 := exec.Command("echo", token)
	cmd2 := exec.Command("terraform", "import", resourceName, id)

	cmd2.Stdin, err = cmd1.StdoutPipe()
	if err != nil {
		return err
	}
	cmd2.Stdout = os.Stdout

	if err = cmd2.Start(); err != nil {
		return err
	}
	if err = cmd1.Run(); err != nil {
		return err
	}
	if err = cmd2.Wait(); err != nil {
		return err
	}

	return nil
}

// importResource forms the data required to import a resource and calls import
// method.
func importResource(token, resType, name, id string) error {
	resourceName, err := getResourceName(resType, name)
	if err != nil {
		return err
	}
	if err := terraformImport(token, resourceName, id); err != nil {
		return err
	}
	return nil
}

// NewTf returns a new TfClient object.
func NewTf(pdToken string) *TfClient {
	return &TfClient{
		pdToken:   pdToken,
		importRes: importResource,
	}
}
