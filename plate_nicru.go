//go:build lego_nicru

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/nicru"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "nicru" {
        return nil, fmt.Errorf("this build of legotapas only supports `nicru` as a provider")
    }

    return nicru.NewDNSProvider()
}
