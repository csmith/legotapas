//go:build lego_regfish

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/regfish"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "regfish" {
        return nil, fmt.Errorf("this build of legotapas only supports `regfish` as a provider")
    }

    return regfish.NewDNSProvider()
}
