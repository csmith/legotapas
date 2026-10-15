//go:build lego_dnsmint

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dnsmint"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dnsmint" {
		return nil, fmt.Errorf("this build of legotapas only supports `dnsmint` as a provider")
	}

	return dnsmint.NewDNSProvider()
}
