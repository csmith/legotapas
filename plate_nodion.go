//go:build lego_nodion

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/nodion"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "nodion" {
        return nil, fmt.Errorf("this build of legotapas only supports `nodion` as a provider")
    }

    return nodion.NewDNSProvider()
}
