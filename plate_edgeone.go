//go:build lego_edgeone

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/edgeone"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "edgeone" {
        return nil, fmt.Errorf("this build of legotapas only supports `edgeone` as a provider")
    }

    return edgeone.NewDNSProvider()
}
