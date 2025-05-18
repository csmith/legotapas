//go:build lego_directadmin

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/directadmin"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "directadmin" {
        return nil, fmt.Errorf("this build of legotapas only supports `directadmin` as a provider")
    }

    return directadmin.NewDNSProvider()
}
