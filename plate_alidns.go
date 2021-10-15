//go:build lego_alidns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/alidns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "alidns" {
        return nil, fmt.Errorf("this build of legotapas only supports `alidns` as a provider")
    }

    return alidns.NewDNSProvider()
}
