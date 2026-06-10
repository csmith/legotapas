//go:build lego_derak

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/derak"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "derak" {
		return nil, fmt.Errorf("this build of legotapas only supports `derak` as a provider")
	}

	return derak.NewDNSProvider()
}
