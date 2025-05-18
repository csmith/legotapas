//go:build lego_westcn

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/westcn"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "westcn" {
        return nil, fmt.Errorf("this build of legotapas only supports `westcn` as a provider")
    }

    return westcn.NewDNSProvider()
}
