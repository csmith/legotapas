//go:build lego_iwantmyname

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/iwantmyname"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "iwantmyname" {
        return nil, fmt.Errorf("this build of legotapas only supports `iwantmyname` as a provider")
    }

    return iwantmyname.NewDNSProvider()
}
