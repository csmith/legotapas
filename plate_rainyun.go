//go:build lego_rainyun

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/rainyun"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "rainyun" {
		return nil, fmt.Errorf("this build of legotapas only supports `rainyun` as a provider")
	}

	return rainyun.NewDNSProvider()
}
