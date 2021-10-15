//go:build lego_cloudflare

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/cloudflare"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "cloudflare" {
        return nil, fmt.Errorf("this build of legotapas only supports `cloudflare` as a provider")
    }

    return cloudflare.NewDNSProvider()
}
