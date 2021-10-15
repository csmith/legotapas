//go:build lego_pdns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/pdns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "pdns" {
        return nil, fmt.Errorf("this build of legotapas only supports `pdns` as a provider")
    }

    return pdns.NewDNSProvider()
}
