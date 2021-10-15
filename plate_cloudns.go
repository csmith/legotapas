//go:build lego_cloudns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/cloudns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "cloudns" {
        return nil, fmt.Errorf("this build of legotapas only supports `cloudns` as a provider")
    }

    return cloudns.NewDNSProvider()
}
