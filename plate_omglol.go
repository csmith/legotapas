//go:build lego_omglol

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/omglol"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "omglol" {
		return nil, fmt.Errorf("this build of legotapas only supports `omglol` as a provider")
	}

	return omglol.NewDNSProvider()
}
