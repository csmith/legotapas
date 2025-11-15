//go:build lego_beget

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/beget"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "beget" {
        return nil, fmt.Errorf("this build of legotapas only supports `beget` as a provider")
    }

    return beget.NewDNSProvider()
}
