//go:build lego_webnamesca

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/webnamesca"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "webnamesca" {
        return nil, fmt.Errorf("this build of legotapas only supports `webnamesca` as a provider")
    }

    return webnamesca.NewDNSProvider()
}
