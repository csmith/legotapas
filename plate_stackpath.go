//go:build lego_stackpath

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/stackpath"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "stackpath" {
        return nil, fmt.Errorf("this build of legotapas only supports `stackpath` as a provider")
    }

    return stackpath.NewDNSProvider()
}
