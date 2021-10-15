//go:build lego_constellix

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/constellix"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "constellix" {
        return nil, fmt.Errorf("this build of legotapas only supports `constellix` as a provider")
    }

    return constellix.NewDNSProvider()
}
