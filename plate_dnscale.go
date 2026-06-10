//go:build lego_dnscale

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dnscale"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dnscale" {
		return nil, fmt.Errorf("this build of legotapas only supports `dnscale` as a provider")
	}

	return dnscale.NewDNSProvider()
}
