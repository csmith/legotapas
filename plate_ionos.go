//go:build lego_ionos

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/ionos"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "ionos" {
        return nil, fmt.Errorf("this build of legotapas only supports `ionos` as a provider")
    }

    return ionos.NewDNSProvider()
}
