//go:build lego_httpnet

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/httpnet"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "httpnet" {
        return nil, fmt.Errorf("this build of legotapas only supports `httpnet` as a provider")
    }

    return httpnet.NewDNSProvider()
}
