//go:build lego_mailinabox

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/mailinabox"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "mailinabox" {
        return nil, fmt.Errorf("this build of legotapas only supports `mailinabox` as a provider")
    }

    return mailinabox.NewDNSProvider()
}
