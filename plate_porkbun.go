//go:build lego_porkbun

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/porkbun"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "porkbun" {
        return nil, fmt.Errorf("this build of legotapas only supports `porkbun` as a provider")
    }

    return porkbun.NewDNSProvider()
}
