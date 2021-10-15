//go:build lego_otc

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/otc"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "otc" {
        return nil, fmt.Errorf("this build of legotapas only supports `otc` as a provider")
    }

    return otc.NewDNSProvider()
}
