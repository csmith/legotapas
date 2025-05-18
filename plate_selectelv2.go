//go:build lego_selectelv2

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/selectelv2"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "selectelv2" {
        return nil, fmt.Errorf("this build of legotapas only supports `selectelv2` as a provider")
    }

    return selectelv2.NewDNSProvider()
}
