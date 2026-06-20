//go:build lego_connbyte

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/connbyte"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "connbyte" {
		return nil, fmt.Errorf("this build of legotapas only supports `connbyte` as a provider")
	}

	return connbyte.NewDNSProvider()
}
