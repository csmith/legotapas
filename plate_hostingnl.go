//go:build lego_hostingnl

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/hostingnl"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "hostingnl" {
        return nil, fmt.Errorf("this build of legotapas only supports `hostingnl` as a provider")
    }

    return hostingnl.NewDNSProvider()
}
