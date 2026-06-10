//go:build lego_tele3

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/tele3"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "tele3" {
		return nil, fmt.Errorf("this build of legotapas only supports `tele3` as a provider")
	}

	return tele3.NewDNSProvider()
}
