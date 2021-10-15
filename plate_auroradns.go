//go:build lego_auroradns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/auroradns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "auroradns" {
        return nil, fmt.Errorf("this build of legotapas only supports `auroradns` as a provider")
    }

    return auroradns.NewDNSProvider()
}
