//go:build lego_technitium

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/technitium"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "technitium" {
        return nil, fmt.Errorf("this build of legotapas only supports `technitium` as a provider")
    }

    return technitium.NewDNSProvider()
}
