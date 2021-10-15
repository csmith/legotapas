//go:build lego_dode

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/dode"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "dode" {
        return nil, fmt.Errorf("this build of legotapas only supports `dode` as a provider")
    }

    return dode.NewDNSProvider()
}
