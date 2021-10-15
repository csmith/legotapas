//go:build lego_loopia

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/loopia"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "loopia" {
        return nil, fmt.Errorf("this build of legotapas only supports `loopia` as a provider")
    }

    return loopia.NewDNSProvider()
}
