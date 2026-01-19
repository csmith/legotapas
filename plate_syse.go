//go:build lego_syse

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/syse"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "syse" {
        return nil, fmt.Errorf("this build of legotapas only supports `syse` as a provider")
    }

    return syse.NewDNSProvider()
}
