//go:build lego_iij

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/iij"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "iij" {
        return nil, fmt.Errorf("this build of legotapas only supports `iij` as a provider")
    }

    return iij.NewDNSProvider()
}
