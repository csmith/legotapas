//go:build lego_ionoscloud

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/ionoscloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "ionoscloud" {
        return nil, fmt.Errorf("this build of legotapas only supports `ionoscloud` as a provider")
    }

    return ionoscloud.NewDNSProvider()
}
