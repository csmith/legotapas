//go:build lego_epik

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/epik"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "epik" {
        return nil, fmt.Errorf("this build of legotapas only supports `epik` as a provider")
    }

    return epik.NewDNSProvider()
}
