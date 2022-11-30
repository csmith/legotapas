//go:build lego_yandexcloud

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/yandexcloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "yandexcloud" {
        return nil, fmt.Errorf("this build of legotapas only supports `yandexcloud` as a provider")
    }

    return yandexcloud.NewDNSProvider()
}
