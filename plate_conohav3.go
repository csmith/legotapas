//go:build lego_conohav3

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/conohav3"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "conohav3" {
        return nil, fmt.Errorf("this build of legotapas only supports `conohav3` as a provider")
    }

    return conohav3.NewDNSProvider()
}
