//go:build lego_gcloud

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/gcloud"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "gcloud" {
        return nil, fmt.Errorf("this build of legotapas only supports `gcloud` as a provider")
    }

    return gcloud.NewDNSProvider()
}
