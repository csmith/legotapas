//go:build lego_gandiv5

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/gandiv5"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "gandiv5" {
		return nil, fmt.Errorf("this build of legotapas only supports `gandiv5` as a provider")
	}

	return gandiv5.NewDNSProvider()
}
