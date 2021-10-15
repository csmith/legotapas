//go:build !lego_cloudns && !lego_bluecat && !lego_selectel && !lego_bindman && !lego_designate && !lego_gcore && !lego_httpreq && !lego_godaddy && !lego_duckdns && !lego_zonomi && !lego_wedos && !lego_simply && !lego_hurricane && !lego_checkdomain && !lego_dynu && !lego_hosttech && !lego_stackpath && !lego_loopia && !lego_infomaniak && !lego_dnsimple && !lego_digitalocean && !lego_dode && !lego_vscale && !lego_dyn && !lego_epik && !lego_namesilo && !lego_nifcloud && !lego_allinkl && !lego_clouddns && !lego_rackspace && !lego_dnspod && !lego_arvancloud && !lego_hostingde && !lego_dnsmadeeasy && !lego_otc && !lego_ionos && !lego_easydns && !lego_ovh && !lego_edgedns && !lego_desec && !lego_liquidweb && !lego_constellix && !lego_joker && !lego_hyperone && !lego_conoha && !lego_vultr && !lego_vinyldns && !lego_gandi && !lego_azure && !lego_exoscale && !lego_mydnsjp && !lego_porkbun && !lego_acmedns && !lego_cloudflare && !lego_route53 && !lego_gcloud && !lego_sakuracloud && !lego_scaleway && !lego_sonic && !lego_nicmanager && !lego_iij && !lego_netlify && !lego_versio && !lego_freemyip && !lego_internetbs && !lego_linode && !lego_zoneee && !lego_domeneshop && !lego_dreamhost && !lego_cloudxns && !lego_luadns && !lego_oraclecloud && !lego_inwx && !lego_ibmcloud && !lego_rfc2136 && !lego_pdns && !lego_vegadns && !lego_lightsail && !lego_yandex && !lego_glesys && !lego_exec && !lego_ns1 && !lego_regru && !lego_servercow && !lego_netcup && !lego_namecheap && !lego_auroradns && !lego_namedotcom && !lego_hetzner && !lego_mythicbeasts && !lego_rimuhosting && !lego_transip && !lego_gandiv5 && !lego_infoblox && !lego_alidns && !lego_autodns && !lego_njalla

package legotapas

import (
    "github.com/go-acme/lego/v4/challenge"
    "github.com/go-acme/lego/v4/providers/dns"
)

func CreateProvider(providerName string) (challenge.Provider, error) {
    return dns.NewDNSChallengeProviderByName(providerName)
}
