//go:build lego_ngenix

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/ngenix"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "ngenix" {
		return nil, fmt.Errorf("this build of legotapas only supports `ngenix` as a provider")
	}

	return ngenix.NewDNSProvider()
}
