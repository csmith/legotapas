//go:build lego_glesys

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/glesys"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "glesys" {
        return nil, fmt.Errorf("this build of legotapas only supports `glesys` as a provider")
    }

    return glesys.NewDNSProvider()
}
