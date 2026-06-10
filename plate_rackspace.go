//go:build lego_rackspace

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/rackspace"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "rackspace" {
		return nil, fmt.Errorf("this build of legotapas only supports `rackspace` as a provider")
	}

	return rackspace.NewDNSProvider()
}
