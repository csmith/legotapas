//go:build lego_myaddr

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/myaddr"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "myaddr" {
        return nil, fmt.Errorf("this build of legotapas only supports `myaddr` as a provider")
    }

    return myaddr.NewDNSProvider()
}
