//go:build lego_route53

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/route53"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "route53" {
        return nil, fmt.Errorf("this build of legotapas only supports `route53` as a provider")
    }

    return route53.NewDNSProvider()
}
