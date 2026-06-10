//go:build lego_hetzner

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/hetzner"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "hetzner" {
		return nil, fmt.Errorf("this build of legotapas only supports `hetzner` as a provider")
	}

	return hetzner.NewDNSProvider()
}
