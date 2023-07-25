//go:build lego_rcodezero

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/rcodezero"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "rcodezero" {
        return nil, fmt.Errorf("this build of legotapas only supports `rcodezero` as a provider")
    }

    return rcodezero.NewDNSProvider()
}
