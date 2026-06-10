//go:build lego_nearlyfreespeech

package legotapas

import (
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/providers/dns/nearlyfreespeech"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
	if strings.ToLower(providerName) != "nearlyfreespeech" {
		return nil, fmt.Errorf("this build of legotapas only supports `nearlyfreespeech` as a provider")
	}

	return nearlyfreespeech.NewDNSProvider()
}
