//go:build lego_internetbs

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/internetbs"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "internetbs" {
		return nil, fmt.Errorf("this build of legotapas only supports `internetbs` as a provider")
	}

	return internetbs.NewDNSProvider()
}
