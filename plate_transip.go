//go:build lego_transip

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/transip"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "transip" {
        return nil, fmt.Errorf("this build of legotapas only supports `transip` as a provider")
    }

    return transip.NewDNSProvider()
}
