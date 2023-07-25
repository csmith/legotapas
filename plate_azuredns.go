//go:build lego_azuredns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/azuredns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "azuredns" {
        return nil, fmt.Errorf("this build of legotapas only supports `azuredns` as a provider")
    }

    return azuredns.NewDNSProvider()
}
