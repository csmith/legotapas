//go:build lego_dnsexit

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/dnsexit"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "dnsexit" {
        return nil, fmt.Errorf("this build of legotapas only supports `dnsexit` as a provider")
    }

    return dnsexit.NewDNSProvider()
}
