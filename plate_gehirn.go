//go:build lego_gehirn

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/gehirn"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "gehirn" {
		return nil, fmt.Errorf("this build of legotapas only supports `gehirn` as a provider")
	}

	return gehirn.NewDNSProvider()
}
