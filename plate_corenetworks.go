//go:build lego_corenetworks

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/corenetworks"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "corenetworks" {
        return nil, fmt.Errorf("this build of legotapas only supports `corenetworks` as a provider")
    }

    return corenetworks.NewDNSProvider()
}
