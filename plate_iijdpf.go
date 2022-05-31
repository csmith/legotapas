//go:build lego_iijdpf

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/iijdpf"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "iijdpf" {
        return nil, fmt.Errorf("this build of legotapas only supports `iijdpf` as a provider")
    }

    return iijdpf.NewDNSProvider()
}
