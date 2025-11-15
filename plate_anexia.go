//go:build lego_anexia

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/anexia"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "anexia" {
        return nil, fmt.Errorf("this build of legotapas only supports `anexia` as a provider")
    }

    return anexia.NewDNSProvider()
}
