//go:build lego_dnspod

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/dnspod"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "dnspod" {
        return nil, fmt.Errorf("this build of legotapas only supports `dnspod` as a provider")
    }

    return dnspod.NewDNSProvider()
}
