//go:build lego_virtualname

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/virtualname"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "virtualname" {
        return nil, fmt.Errorf("this build of legotapas only supports `virtualname` as a provider")
    }

    return virtualname.NewDNSProvider()
}
