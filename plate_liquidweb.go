//go:build lego_liquidweb

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/liquidweb"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "liquidweb" {
        return nil, fmt.Errorf("this build of legotapas only supports `liquidweb` as a provider")
    }

    return liquidweb.NewDNSProvider()
}
