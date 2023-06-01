//go:build lego_googledomains

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/googledomains"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "googledomains" {
        return nil, fmt.Errorf("this build of legotapas only supports `googledomains` as a provider")
    }

    return googledomains.NewDNSProvider()
}
