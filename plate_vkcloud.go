//go:build lego_vkcloud

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/vkcloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "vkcloud" {
        return nil, fmt.Errorf("this build of legotapas only supports `vkcloud` as a provider")
    }

    return vkcloud.NewDNSProvider()
}
