//go:build lego_ucloud

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/ucloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "ucloud" {
        return nil, fmt.Errorf("this build of legotapas only supports `ucloud` as a provider")
    }

    return ucloud.NewDNSProvider()
}
