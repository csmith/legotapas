//go:build lego_regru

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/regru"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "regru" {
        return nil, fmt.Errorf("this build of legotapas only supports `regru` as a provider")
    }

    return regru.NewDNSProvider()
}
