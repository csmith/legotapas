//go:build lego_dnshomede

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/dnshomede"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "dnshomede" {
        return nil, fmt.Errorf("this build of legotapas only supports `dnshomede` as a provider")
    }

    return dnshomede.NewDNSProvider()
}
