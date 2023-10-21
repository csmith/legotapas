//go:build lego_yandex360

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/yandex360"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "yandex360" {
        return nil, fmt.Errorf("this build of legotapas only supports `yandex360` as a provider")
    }

    return yandex360.NewDNSProvider()
}
