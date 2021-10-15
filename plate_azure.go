//go:build lego_azure

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/azure"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "azure" {
        return nil, fmt.Errorf("this build of legotapas only supports `azure` as a provider")
    }

    return azure.NewDNSProvider()
}
