//go:build lego_manual

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/manual"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "manual" {
        return nil, fmt.Errorf("this build of legotapas only supports `manual` as a provider")
    }

    return manual.NewDNSProvider()
}
