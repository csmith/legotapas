//go:build lego_netcup

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/netcup"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "netcup" {
        return nil, fmt.Errorf("this build of legotapas only supports `netcup` as a provider")
    }

    return netcup.NewDNSProvider()
}
