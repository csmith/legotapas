//go:build lego_dnsla

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dnsla"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dnsla" {
		return nil, fmt.Errorf("this build of legotapas only supports `dnsla` as a provider")
	}

	return dnsla.NewDNSProvider()
}
