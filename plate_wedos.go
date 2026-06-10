//go:build lego_wedos

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/wedos"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "wedos" {
		return nil, fmt.Errorf("this build of legotapas only supports `wedos` as a provider")
	}

	return wedos.NewDNSProvider()
}
