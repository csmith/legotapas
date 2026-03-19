//go:build lego_todaynic

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/todaynic"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "todaynic" {
        return nil, fmt.Errorf("this build of legotapas only supports `todaynic` as a provider")
    }

    return todaynic.NewDNSProvider()
}
