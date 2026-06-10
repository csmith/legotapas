//go:build lego_hostingde

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/hostingde"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "hostingde" {
		return nil, fmt.Errorf("this build of legotapas only supports `hostingde` as a provider")
	}

	return hostingde.NewDNSProvider()
}
