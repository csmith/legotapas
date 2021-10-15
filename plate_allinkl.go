//go:build lego_allinkl

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/allinkl"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "allinkl" {
        return nil, fmt.Errorf("this build of legotapas only supports `allinkl` as a provider")
    }

    return allinkl.NewDNSProvider()
}
