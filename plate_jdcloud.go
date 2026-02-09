//go:build lego_jdcloud

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/jdcloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "jdcloud" {
        return nil, fmt.Errorf("this build of legotapas only supports `jdcloud` as a provider")
    }

    return jdcloud.NewDNSProvider()
}
