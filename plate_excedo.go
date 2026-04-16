//go:build lego_excedo

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/excedo"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "excedo" {
        return nil, fmt.Errorf("this build of legotapas only supports `excedo` as a provider")
    }

    return excedo.NewDNSProvider()
}
