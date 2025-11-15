//go:build lego_octenium

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/octenium"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "octenium" {
        return nil, fmt.Errorf("this build of legotapas only supports `octenium` as a provider")
    }

    return octenium.NewDNSProvider()
}
