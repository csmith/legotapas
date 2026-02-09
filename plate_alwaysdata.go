//go:build lego_alwaysdata

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/alwaysdata"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "alwaysdata" {
        return nil, fmt.Errorf("this build of legotapas only supports `alwaysdata` as a provider")
    }

    return alwaysdata.NewDNSProvider()
}
