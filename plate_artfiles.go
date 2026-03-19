//go:build lego_artfiles

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/artfiles"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "artfiles" {
        return nil, fmt.Errorf("this build of legotapas only supports `artfiles` as a provider")
    }

    return artfiles.NewDNSProvider()
}
