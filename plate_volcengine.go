//go:build lego_volcengine

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/volcengine"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "volcengine" {
        return nil, fmt.Errorf("this build of legotapas only supports `volcengine` as a provider")
    }

    return volcengine.NewDNSProvider()
}
