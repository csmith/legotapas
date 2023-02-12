//go:build lego_websupport

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/websupport"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "websupport" {
        return nil, fmt.Errorf("this build of legotapas only supports `websupport` as a provider")
    }

    return websupport.NewDNSProvider()
}
