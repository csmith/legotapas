//go:build lego_desec

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/desec"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "desec" {
        return nil, fmt.Errorf("this build of legotapas only supports `desec` as a provider")
    }

    return desec.NewDNSProvider()
}
