//go:build lego_poweradmin

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/poweradmin"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "poweradmin" {
		return nil, fmt.Errorf("this build of legotapas only supports `poweradmin` as a provider")
	}

	return poweradmin.NewDNSProvider()
}
