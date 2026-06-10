//go:build lego_rage4

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/rage4"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "rage4" {
		return nil, fmt.Errorf("this build of legotapas only supports `rage4` as a provider")
	}

	return rage4.NewDNSProvider()
}
