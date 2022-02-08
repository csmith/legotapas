//go:build lego_safedns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/safedns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "safedns" {
        return nil, fmt.Errorf("this build of legotapas only supports `safedns` as a provider")
    }

    return safedns.NewDNSProvider()
}
