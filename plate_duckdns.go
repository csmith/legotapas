//go:build lego_duckdns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/duckdns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "duckdns" {
        return nil, fmt.Errorf("this build of legotapas only supports `duckdns` as a provider")
    }

    return duckdns.NewDNSProvider()
}
