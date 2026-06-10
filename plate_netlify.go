//go:build lego_netlify

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/netlify"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "netlify" {
		return nil, fmt.Errorf("this build of legotapas only supports `netlify` as a provider")
	}

	return netlify.NewDNSProvider()
}
