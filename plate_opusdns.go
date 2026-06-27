//go:build lego_opusdns

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/opusdns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "opusdns" {
		return nil, fmt.Errorf("this build of legotapas only supports `opusdns` as a provider")
	}

	return opusdns.NewDNSProvider()
}
