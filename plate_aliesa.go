//go:build lego_aliesa

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/aliesa"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "aliesa" {
        return nil, fmt.Errorf("this build of legotapas only supports `aliesa` as a provider")
    }

    return aliesa.NewDNSProvider()
}
