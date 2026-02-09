//go:build lego_ispconfigddns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/ispconfigddns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "ispconfigddns" {
        return nil, fmt.Errorf("this build of legotapas only supports `ispconfigddns` as a provider")
    }

    return ispconfigddns.NewDNSProvider()
}
