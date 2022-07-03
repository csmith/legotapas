//go:build lego_variomedia

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/variomedia"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "variomedia" {
        return nil, fmt.Errorf("this build of legotapas only supports `variomedia` as a provider")
    }

    return variomedia.NewDNSProvider()
}
