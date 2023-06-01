//go:build lego_brandit

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/brandit"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "brandit" {
        return nil, fmt.Errorf("this build of legotapas only supports `brandit` as a provider")
    }

    return brandit.NewDNSProvider()
}
