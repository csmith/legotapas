//go:build lego_neodigit

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/neodigit"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "neodigit" {
        return nil, fmt.Errorf("this build of legotapas only supports `neodigit` as a provider")
    }

    return neodigit.NewDNSProvider()
}
