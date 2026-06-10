//go:build lego_veesp

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/veesp"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "veesp" {
		return nil, fmt.Errorf("this build of legotapas only supports `veesp` as a provider")
	}

	return veesp.NewDNSProvider()
}
