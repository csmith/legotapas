//go:build lego_infoblox

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/infoblox"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "infoblox" {
		return nil, fmt.Errorf("this build of legotapas only supports `infoblox` as a provider")
	}

	return infoblox.NewDNSProvider()
}
