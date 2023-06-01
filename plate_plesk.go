//go:build lego_plesk

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/plesk"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "plesk" {
        return nil, fmt.Errorf("this build of legotapas only supports `plesk` as a provider")
    }

    return plesk.NewDNSProvider()
}
