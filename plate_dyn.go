//go:build lego_dyn

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dyn"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dyn" {
		return nil, fmt.Errorf("this build of legotapas only supports `dyn` as a provider")
	}

	return dyn.NewDNSProvider()
}
