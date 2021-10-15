//go:build lego_dnsmadeeasy

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/dnsmadeeasy"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "dnsmadeeasy" {
        return nil, fmt.Errorf("this build of legotapas only supports `dnsmadeeasy` as a provider")
    }

    return dnsmadeeasy.NewDNSProvider()
}
