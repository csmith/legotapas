//go:build lego_versio

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/versio"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "versio" {
		return nil, fmt.Errorf("this build of legotapas only supports `versio` as a provider")
	}

	return versio.NewDNSProvider()
}
