//go:build lego_vercel

package legotapas

import (
    "fmt"
    "strings"

    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns/vercel"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    if strings.ToLower(providerName) != "vercel" {
        return nil, fmt.Errorf("this build of legotapas only supports `vercel` as a provider")
    }

    return vercel.NewDNSProvider()
}
