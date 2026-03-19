//go:build lego_ddnss

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/ddnss"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "ddnss" {
        return nil, fmt.Errorf("this build of legotapas only supports `ddnss` as a provider")
    }

    return ddnss.NewDNSProvider()
}
