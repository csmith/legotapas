//go:build lego_scaleway

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/scaleway"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "scaleway" {
        return nil, fmt.Errorf("this build of legotapas only supports `scaleway` as a provider")
    }

    return scaleway.NewDNSProvider()
}
