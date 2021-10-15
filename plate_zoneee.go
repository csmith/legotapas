//go:build lego_zoneee

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/zoneee"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "zoneee" {
        return nil, fmt.Errorf("this build of legotapas only supports `zoneee` as a provider")
    }

    return zoneee.NewDNSProvider()
}
