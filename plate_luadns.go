//go:build lego_luadns

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/luadns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "luadns" {
		return nil, fmt.Errorf("this build of legotapas only supports `luadns` as a provider")
	}

	return luadns.NewDNSProvider()
}
