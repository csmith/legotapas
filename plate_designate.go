//go:build lego_designate

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/designate"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "designate" {
        return nil, fmt.Errorf("this build of legotapas only supports `designate` as a provider")
    }

    return designate.NewDNSProvider()
}
