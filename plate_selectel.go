//go:build lego_selectel

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/selectel"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "selectel" {
		return nil, fmt.Errorf("this build of legotapas only supports `selectel` as a provider")
	}

	return selectel.NewDNSProvider()
}
