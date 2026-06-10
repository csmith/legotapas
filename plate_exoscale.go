//go:build lego_exoscale

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/exoscale"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "exoscale" {
		return nil, fmt.Errorf("this build of legotapas only supports `exoscale` as a provider")
	}

	return exoscale.NewDNSProvider()
}
