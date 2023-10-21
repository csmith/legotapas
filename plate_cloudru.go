//go:build lego_cloudru

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/cloudru"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "cloudru" {
        return nil, fmt.Errorf("this build of legotapas only supports `cloudru` as a provider")
    }

    return cloudru.NewDNSProvider()
}
