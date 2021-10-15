//go:build lego_servercow

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/servercow"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "servercow" {
        return nil, fmt.Errorf("this build of legotapas only supports `servercow` as a provider")
    }

    return servercow.NewDNSProvider()
}
