//go:build lego_uniteddomains

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/uniteddomains"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "uniteddomains" {
        return nil, fmt.Errorf("this build of legotapas only supports `uniteddomains` as a provider")
    }

    return uniteddomains.NewDNSProvider()
}
