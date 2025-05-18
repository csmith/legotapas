//go:build lego_f5xc

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/f5xc"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "f5xc" {
        return nil, fmt.Errorf("this build of legotapas only supports `f5xc` as a provider")
    }

    return f5xc.NewDNSProvider()
}
