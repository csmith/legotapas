//go:build lego_euserv

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/euserv"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "euserv" {
		return nil, fmt.Errorf("this build of legotapas only supports `euserv` as a provider")
	}

	return euserv.NewDNSProvider()
}
