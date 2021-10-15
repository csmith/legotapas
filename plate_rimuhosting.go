//go:build lego_rimuhosting

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/rimuhosting"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "rimuhosting" {
        return nil, fmt.Errorf("this build of legotapas only supports `rimuhosting` as a provider")
    }

    return rimuhosting.NewDNSProvider()
}
