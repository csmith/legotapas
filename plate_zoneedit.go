//go:build lego_zoneedit

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/zoneedit"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "zoneedit" {
        return nil, fmt.Errorf("this build of legotapas only supports `zoneedit` as a provider")
    }

    return zoneedit.NewDNSProvider()
}
