//go:build lego_domeneshop

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/domeneshop"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "domeneshop" {
		return nil, fmt.Errorf("this build of legotapas only supports `domeneshop` as a provider")
	}

	return domeneshop.NewDNSProvider()
}
