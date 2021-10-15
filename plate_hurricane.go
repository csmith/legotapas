//go:build lego_hurricane

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/hurricane"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "hurricane" {
        return nil, fmt.Errorf("this build of legotapas only supports `hurricane` as a provider")
    }

    return hurricane.NewDNSProvider()
}
