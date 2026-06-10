//go:build lego_godaddy

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/godaddy"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "godaddy" {
		return nil, fmt.Errorf("this build of legotapas only supports `godaddy` as a provider")
	}

	return godaddy.NewDNSProvider()
}
