//go:build lego_webnames

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/webnames"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "webnames" {
        return nil, fmt.Errorf("this build of legotapas only supports `webnames` as a provider")
    }

    return webnames.NewDNSProvider()
}
