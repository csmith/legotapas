//go:build lego_infomaniak

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/infomaniak"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "infomaniak" {
        return nil, fmt.Errorf("this build of legotapas only supports `infomaniak` as a provider")
    }

    return infomaniak.NewDNSProvider()
}
