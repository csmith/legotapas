//go:build lego_myra

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/myra"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "myra" {
		return nil, fmt.Errorf("this build of legotapas only supports `myra` as a provider")
	}

	return myra.NewDNSProvider()
}
