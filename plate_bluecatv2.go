//go:build lego_bluecatv2

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/bluecatv2"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "bluecatv2" {
        return nil, fmt.Errorf("this build of legotapas only supports `bluecatv2` as a provider")
    }

    return bluecatv2.NewDNSProvider()
}
