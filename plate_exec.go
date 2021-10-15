//go:build lego_exec

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/exec"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "exec" {
        return nil, fmt.Errorf("this build of legotapas only supports `exec` as a provider")
    }

    return exec.NewDNSProvider()
}
