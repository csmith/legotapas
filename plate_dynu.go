//go:build lego_dynu

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/dynu"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "dynu" {
        return nil, fmt.Errorf("this build of legotapas only supports `dynu` as a provider")
    }

    return dynu.NewDNSProvider()
}
