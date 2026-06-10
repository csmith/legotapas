//go:build lego_dinahosting

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/dinahosting"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "dinahosting" {
		return nil, fmt.Errorf("this build of legotapas only supports `dinahosting` as a provider")
	}

	return dinahosting.NewDNSProvider()
}
