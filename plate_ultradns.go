//go:build lego_ultradns

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/ultradns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "ultradns" {
		return nil, fmt.Errorf("this build of legotapas only supports `ultradns` as a provider")
	}

	return ultradns.NewDNSProvider()
}
