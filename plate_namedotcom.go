//go:build lego_namedotcom

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/namedotcom"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "namedotcom" {
		return nil, fmt.Errorf("this build of legotapas only supports `namedotcom` as a provider")
	}

	return namedotcom.NewDNSProvider()
}
