//go:build lego_limacity

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/limacity"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "limacity" {
        return nil, fmt.Errorf("this build of legotapas only supports `limacity` as a provider")
    }

    return limacity.NewDNSProvider()
}
