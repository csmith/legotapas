//go:build lego_cloudxns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/cloudxns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "cloudxns" {
        return nil, fmt.Errorf("this build of legotapas only supports `cloudxns` as a provider")
    }

    return cloudxns.NewDNSProvider()
}
