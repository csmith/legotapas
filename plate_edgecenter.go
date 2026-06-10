//go:build lego_edgecenter

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/edgecenter"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "edgecenter" {
		return nil, fmt.Errorf("this build of legotapas only supports `edgecenter` as a provider")
	}

	return edgecenter.NewDNSProvider()
}
