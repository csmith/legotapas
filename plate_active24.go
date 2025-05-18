//go:build lego_active24

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/active24"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "active24" {
        return nil, fmt.Errorf("this build of legotapas only supports `active24` as a provider")
    }

    return active24.NewDNSProvider()
}
