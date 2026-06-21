//go:build lego_dynadot

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dynadot"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dynadot" {
		return nil, fmt.Errorf("this build of legotapas only supports `dynadot` as a provider")
	}

	return dynadot.NewDNSProvider()
}
