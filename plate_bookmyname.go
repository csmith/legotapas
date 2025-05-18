//go:build lego_bookmyname

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/bookmyname"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "bookmyname" {
        return nil, fmt.Errorf("this build of legotapas only supports `bookmyname` as a provider")
    }

    return bookmyname.NewDNSProvider()
}
