//go:build lego_liara

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/liara"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "liara" {
        return nil, fmt.Errorf("this build of legotapas only supports `liara` as a provider")
    }

    return liara.NewDNSProvider()
}
