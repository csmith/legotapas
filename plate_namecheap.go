//go:build lego_namecheap

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/namecheap"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "namecheap" {
        return nil, fmt.Errorf("this build of legotapas only supports `namecheap` as a provider")
    }

    return namecheap.NewDNSProvider()
}
