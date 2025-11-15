//go:build lego_binarylane

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/binarylane"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "binarylane" {
        return nil, fmt.Errorf("this build of legotapas only supports `binarylane` as a provider")
    }

    return binarylane.NewDNSProvider()
}
