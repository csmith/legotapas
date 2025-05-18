//go:build lego_timewebcloud

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/timewebcloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "timewebcloud" {
        return nil, fmt.Errorf("this build of legotapas only supports `timewebcloud` as a provider")
    }

    return timewebcloud.NewDNSProvider()
}
