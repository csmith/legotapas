//go:build lego_czechia

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/czechia"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "czechia" {
        return nil, fmt.Errorf("this build of legotapas only supports `czechia` as a provider")
    }

    return czechia.NewDNSProvider()
}
