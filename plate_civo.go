//go:build lego_civo

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/civo"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "civo" {
		return nil, fmt.Errorf("this build of legotapas only supports `civo` as a provider")
	}

	return civo.NewDNSProvider()
}
