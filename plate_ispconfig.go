//go:build lego_ispconfig

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/ispconfig"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "ispconfig" {
        return nil, fmt.Errorf("this build of legotapas only supports `ispconfig` as a provider")
    }

    return ispconfig.NewDNSProvider()
}
