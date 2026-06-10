//go:build lego_efficientip

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/efficientip"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "efficientip" {
		return nil, fmt.Errorf("this build of legotapas only supports `efficientip` as a provider")
	}

	return efficientip.NewDNSProvider()
}
