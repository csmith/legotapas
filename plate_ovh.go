//go:build lego_ovh

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/ovh"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "ovh" {
		return nil, fmt.Errorf("this build of legotapas only supports `ovh` as a provider")
	}

	return ovh.NewDNSProvider()
}
