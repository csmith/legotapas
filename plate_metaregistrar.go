//go:build lego_metaregistrar

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/metaregistrar"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "metaregistrar" {
		return nil, fmt.Errorf("this build of legotapas only supports `metaregistrar` as a provider")
	}

	return metaregistrar.NewDNSProvider()
}
