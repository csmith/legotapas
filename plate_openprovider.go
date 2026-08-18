//go:build lego_openprovider

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/openprovider"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "openprovider" {
		return nil, fmt.Errorf("this build of legotapas only supports `openprovider` as a provider")
	}

	return openprovider.NewDNSProvider()
}
