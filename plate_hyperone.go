//go:build lego_hyperone

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/hyperone"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "hyperone" {
        return nil, fmt.Errorf("this build of legotapas only supports `hyperone` as a provider")
    }

    return hyperone.NewDNSProvider()
}
