//go:build lego_sonic

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/sonic"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "sonic" {
        return nil, fmt.Errorf("this build of legotapas only supports `sonic` as a provider")
    }

    return sonic.NewDNSProvider()
}
