//go:build lego_autodns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/autodns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "autodns" {
        return nil, fmt.Errorf("this build of legotapas only supports `autodns` as a provider")
    }

    return autodns.NewDNSProvider()
}
