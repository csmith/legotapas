//go:build lego_inwx

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/inwx"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "inwx" {
        return nil, fmt.Errorf("this build of legotapas only supports `inwx` as a provider")
    }

    return inwx.NewDNSProvider()
}
