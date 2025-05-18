//go:build lego_axelname

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/axelname"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "axelname" {
        return nil, fmt.Errorf("this build of legotapas only supports `axelname` as a provider")
    }

    return axelname.NewDNSProvider()
}
