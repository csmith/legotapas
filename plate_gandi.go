//go:build lego_gandi

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/gandi"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "gandi" {
        return nil, fmt.Errorf("this build of legotapas only supports `gandi` as a provider")
    }

    return gandi.NewDNSProvider()
}
