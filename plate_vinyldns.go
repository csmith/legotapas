//go:build lego_vinyldns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/vinyldns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "vinyldns" {
        return nil, fmt.Errorf("this build of legotapas only supports `vinyldns` as a provider")
    }

    return vinyldns.NewDNSProvider()
}
