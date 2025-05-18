//go:build lego_cpanel

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/cpanel"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "cpanel" {
        return nil, fmt.Errorf("this build of legotapas only supports `cpanel` as a provider")
    }

    return cpanel.NewDNSProvider()
}
