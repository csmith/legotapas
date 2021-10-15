//go:build lego_httpreq

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/httpreq"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "httpreq" {
        return nil, fmt.Errorf("this build of legotapas only supports `httpreq` as a provider")
    }

    return httpreq.NewDNSProvider()
}
