//go:build lego_dyndnsfree

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/dyndnsfree"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "dyndnsfree" {
        return nil, fmt.Errorf("this build of legotapas only supports `dyndnsfree` as a provider")
    }

    return dyndnsfree.NewDNSProvider()
}
