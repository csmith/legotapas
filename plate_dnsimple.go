//go:build lego_dnsimple

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/dnsimple"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "dnsimple" {
        return nil, fmt.Errorf("this build of legotapas only supports `dnsimple` as a provider")
    }

    return dnsimple.NewDNSProvider()
}
