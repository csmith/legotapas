//go:build lego_baiducloud

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/baiducloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "baiducloud" {
        return nil, fmt.Errorf("this build of legotapas only supports `baiducloud` as a provider")
    }

    return baiducloud.NewDNSProvider()
}
