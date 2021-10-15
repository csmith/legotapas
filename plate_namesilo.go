//go:build lego_namesilo

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/namesilo"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "namesilo" {
        return nil, fmt.Errorf("this build of legotapas only supports `namesilo` as a provider")
    }

    return namesilo.NewDNSProvider()
}
