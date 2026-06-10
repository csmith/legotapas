//go:build lego_checkdomain

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/checkdomain"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "checkdomain" {
		return nil, fmt.Errorf("this build of legotapas only supports `checkdomain` as a provider")
	}

	return checkdomain.NewDNSProvider()
}
