//go:build lego_clouddns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/clouddns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "clouddns" {
        return nil, fmt.Errorf("this build of legotapas only supports `clouddns` as a provider")
    }

    return clouddns.NewDNSProvider()
}
