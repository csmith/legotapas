//go:build lego_ipv64

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/ipv64"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "ipv64" {
        return nil, fmt.Errorf("this build of legotapas only supports `ipv64` as a provider")
    }

    return ipv64.NewDNSProvider()
}
