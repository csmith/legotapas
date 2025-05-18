//go:build lego_shellrent

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/shellrent"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "shellrent" {
        return nil, fmt.Errorf("this build of legotapas only supports `shellrent` as a provider")
    }

    return shellrent.NewDNSProvider()
}
