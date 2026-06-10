//go:build lego_dandomain

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dandomain"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dandomain" {
		return nil, fmt.Errorf("this build of legotapas only supports `dandomain` as a provider")
	}

	return dandomain.NewDNSProvider()
}
