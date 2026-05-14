//go:build lego_netnod

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/netnod"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "netnod" {
        return nil, fmt.Errorf("this build of legotapas only supports `netnod` as a provider")
    }

    return netnod.NewDNSProvider()
}
