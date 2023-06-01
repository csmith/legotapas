//go:build lego_bunny

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/bunny"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "bunny" {
        return nil, fmt.Errorf("this build of legotapas only supports `bunny` as a provider")
    }

    return bunny.NewDNSProvider()
}
