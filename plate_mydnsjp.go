//go:build lego_mydnsjp

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/mydnsjp"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "mydnsjp" {
        return nil, fmt.Errorf("this build of legotapas only supports `mydnsjp` as a provider")
    }

    return mydnsjp.NewDNSProvider()
}
