//go:build lego_vscale

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/vscale"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "vscale" {
        return nil, fmt.Errorf("this build of legotapas only supports `vscale` as a provider")
    }

    return vscale.NewDNSProvider()
}
