//go:build lego_leaseweb

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/leaseweb"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "leaseweb" {
        return nil, fmt.Errorf("this build of legotapas only supports `leaseweb` as a provider")
    }

    return leaseweb.NewDNSProvider()
}
