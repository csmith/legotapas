//go:build lego_linode

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/linode"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "linode" {
        return nil, fmt.Errorf("this build of legotapas only supports `linode` as a provider")
    }

    return linode.NewDNSProvider()
}
