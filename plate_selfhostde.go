//go:build lego_selfhostde

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/selfhostde"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "selfhostde" {
        return nil, fmt.Errorf("this build of legotapas only supports `selfhostde` as a provider")
    }

    return selfhostde.NewDNSProvider()
}
