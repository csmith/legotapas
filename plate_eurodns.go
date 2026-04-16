//go:build lego_eurodns

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/eurodns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "eurodns" {
        return nil, fmt.Errorf("this build of legotapas only supports `eurodns` as a provider")
    }

    return eurodns.NewDNSProvider()
}
