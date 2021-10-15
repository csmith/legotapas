//go:build lego_bindman

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/bindman"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "bindman" {
        return nil, fmt.Errorf("this build of legotapas only supports `bindman` as a provider")
    }

    return bindman.NewDNSProvider()
}
