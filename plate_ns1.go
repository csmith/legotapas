//go:build lego_ns1

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/ns1"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "ns1" {
        return nil, fmt.Errorf("this build of legotapas only supports `ns1` as a provider")
    }

    return ns1.NewDNSProvider()
}
