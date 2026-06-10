//go:build lego_gname

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/gname"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "gname" {
		return nil, fmt.Errorf("this build of legotapas only supports `gname` as a provider")
	}

	return gname.NewDNSProvider()
}
