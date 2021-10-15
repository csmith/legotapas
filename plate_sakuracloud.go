//go:build lego_sakuracloud

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/sakuracloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "sakuracloud" {
        return nil, fmt.Errorf("this build of legotapas only supports `sakuracloud` as a provider")
    }

    return sakuracloud.NewDNSProvider()
}
