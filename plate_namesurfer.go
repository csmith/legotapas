//go:build lego_namesurfer

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/namesurfer"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "namesurfer" {
        return nil, fmt.Errorf("this build of legotapas only supports `namesurfer` as a provider")
    }

    return namesurfer.NewDNSProvider()
}
