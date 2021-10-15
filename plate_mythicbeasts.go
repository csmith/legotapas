//go:build lego_mythicbeasts

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/mythicbeasts"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "mythicbeasts" {
        return nil, fmt.Errorf("this build of legotapas only supports `mythicbeasts` as a provider")
    }

    return mythicbeasts.NewDNSProvider()
}
