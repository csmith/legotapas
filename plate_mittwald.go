//go:build lego_mittwald

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/mittwald"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "mittwald" {
        return nil, fmt.Errorf("this build of legotapas only supports `mittwald` as a provider")
    }

    return mittwald.NewDNSProvider()
}
