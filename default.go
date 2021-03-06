//go:build !lego_acmedns && !lego_alidns && !lego_allinkl && !lego_arvancloud && !lego_auroradns && !lego_autodns && !lego_azure && !lego_bindman && !lego_bluecat && !lego_checkdomain && !lego_clouddns && !lego_cloudflare && !lego_cloudns && !lego_cloudxns && !lego_conoha && !lego_constellix && !lego_desec && !lego_designate && !lego_digitalocean && !lego_dnsimple && !lego_dnsmadeeasy && !lego_dnspod && !lego_dode && !lego_domeneshop && !lego_dreamhost && !lego_duckdns && !lego_dyn && !lego_dynu && !lego_easydns && !lego_edgedns && !lego_epik && !lego_exec && !lego_exoscale && !lego_freemyip && !lego_gandi && !lego_gandiv5 && !lego_gcloud && !lego_gcore && !lego_glesys && !lego_godaddy && !lego_hetzner && !lego_hostingde && !lego_hosttech && !lego_httpreq && !lego_hurricane && !lego_hyperone && !lego_ibmcloud && !lego_iij && !lego_iijdpf && !lego_infoblox && !lego_infomaniak && !lego_internetbs && !lego_inwx && !lego_ionos && !lego_iwantmyname && !lego_joker && !lego_lightsail && !lego_linode && !lego_liquidweb && !lego_loopia && !lego_luadns && !lego_mydnsjp && !lego_mythicbeasts && !lego_namecheap && !lego_namedotcom && !lego_namesilo && !lego_nearlyfreespeech && !lego_netcup && !lego_netlify && !lego_nicmanager && !lego_nifcloud && !lego_njalla && !lego_ns1 && !lego_oraclecloud && !lego_otc && !lego_ovh && !lego_pdns && !lego_porkbun && !lego_rackspace && !lego_regru && !lego_rfc2136 && !lego_rimuhosting && !lego_route53 && !lego_safedns && !lego_sakuracloud && !lego_scaleway && !lego_selectel && !lego_servercow && !lego_simply && !lego_sonic && !lego_stackpath && !lego_tencentcloud && !lego_transip && !lego_variomedia && !lego_vegadns && !lego_vercel && !lego_versio && !lego_vinyldns && !lego_vscale && !lego_vultr && !lego_wedos && !lego_yandex && !lego_zoneee && !lego_zonomi

package legotapas

import (
    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    return dns.NewDNSChallengeProviderByName(providerName)
}
