//go:build lego_gigahostno

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/gigahostno"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "gigahostno" {
		return nil, fmt.Errorf("this build of legotapas only supports `gigahostno` as a provider")
	}

	return gigahostno.NewDNSProvider()
}
