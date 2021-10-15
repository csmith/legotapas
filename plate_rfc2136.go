//go:build lego_rfc2136

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/rfc2136"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "rfc2136" {
        return nil, fmt.Errorf("this build of legotapas only supports `rfc2136` as a provider")
    }

    return rfc2136.NewDNSProvider()
}
