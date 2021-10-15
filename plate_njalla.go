//go:build lego_njalla

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/njalla"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "njalla" {
        return nil, fmt.Errorf("this build of legotapas only supports `njalla` as a provider")
    }

    return njalla.NewDNSProvider()
}
