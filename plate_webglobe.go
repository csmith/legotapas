//go:build lego_webglobe

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/webglobe"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "webglobe" {
		return nil, fmt.Errorf("this build of legotapas only supports `webglobe` as a provider")
	}

	return webglobe.NewDNSProvider()
}
