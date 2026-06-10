//go:build lego_dnsservices

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dnsservices"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dnsservices" {
		return nil, fmt.Errorf("this build of legotapas only supports `dnsservices` as a provider")
	}

	return dnsservices.NewDNSProvider()
}
