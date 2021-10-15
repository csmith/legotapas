//go:build lego_yandex

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/yandex"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "yandex" {
        return nil, fmt.Errorf("this build of legotapas only supports `yandex` as a provider")
    }

    return yandex.NewDNSProvider()
}
