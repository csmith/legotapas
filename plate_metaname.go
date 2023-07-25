//go:build lego_metaname

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/metaname"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "metaname" {
        return nil, fmt.Errorf("this build of legotapas only supports `metaname` as a provider")
    }

    return metaname.NewDNSProvider()
}
