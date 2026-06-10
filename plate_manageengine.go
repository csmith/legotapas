//go:build lego_manageengine

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/manageengine"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "manageengine" {
		return nil, fmt.Errorf("this build of legotapas only supports `manageengine` as a provider")
	}

	return manageengine.NewDNSProvider()
}
