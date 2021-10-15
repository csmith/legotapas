//go:build lego_hosttech

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/hosttech"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "hosttech" {
        return nil, fmt.Errorf("this build of legotapas only supports `hosttech` as a provider")
    }

    return hosttech.NewDNSProvider()
}
