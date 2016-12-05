package main

import (
    // "fmt"
    "strings"
    "os"
    "os/exec"
    "errors"

    "github.com/spf13/viper"
)

const (
    EscalationPolicyPrefix = "pagerduty_escalation_policy"
    ServicePrefix = "pagerduty_service"
)

type TerraformClient interface {
    ImportService(id, name string)
    ImportEscalationPolicy(id, name string)
}

type TfClient struct {
    pdToken string
    importRes func(token, resType, name, id string) error
}

func (c TfClient) ImportEscalationPolicy(id, name string) error {
    return c.importRes(c.pdToken, EscalationPolicyPrefix, name, id)
}

func (c TfClient) ImportService(id, name string) error {
    return c.importRes(c.pdToken, ServicePrefix, name, id)
}

func getResourceName(resType, name string) (string, error) {
    switch resType {
    case EscalationPolicyPrefix:
        return strings.Join([]string{EscalationPolicyPrefix, name}, "."), nil
    case ServicePrefix:
        return strings.Join([]string{ServicePrefix, name}, "."), nil
    default:
        return "", errors.New("Unknown resource type" + name)
    }
}

func importResource(token, resType, name, id string) error {
    resourceName, err := getResourceName(resType, name)
    if err != nil {
        return err
    }
    cmd1 := exec.Command("echo", token)
    cmd2 := exec.Command("terraform", "import", resourceName, id)

    // Pipe output of cmd1 as input of cmd2.
    // This is required because `terraform import` doesn't allow passing
    // variable arguments in commandline at the moment.
    cmd2.Stdin, _ = cmd1.StdoutPipe()
    cmd2.Stdout = os.Stdout

    cmd2.Start()
    cmd1.Run()
    cmd2.Wait()

    return nil
}

func main() {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }
    authtoken := viper.GetString("authtoken")
    tfclient := TfClient{pdToken: authtoken, importRes: importResource}
    tfclient.ImportService("P595V8T", "acl")
}
