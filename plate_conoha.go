//go:build lego_conoha

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/conoha"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "conoha" {
		return nil, fmt.Errorf("this build of legotapas only supports `conoha` as a provider")
	}

	return conoha.NewDNSProvider()
}
