# Lego Tapas

Provides build constraints for the DNS providers found in [lego](https://github.com/go-acme/lego).

## How to use

Normally when using lego you'd call `dns.NewDNSChallengeProviderByName(providerName string)` to get
an instance of the DNS provider. To use legotapas replace this with a call to
`legotapas.CreateProvider(providerName string)`. It has the same signature, so shouldn't take more
than a couple of seconds.  If you don't specify any build constraints then this will function exactly
as it did before.

Now you can use build constraints to select a single provider to compile with, instead of pulling
in every single provider and all of their dependencies. For example to support only the `httpreq`
provider you would build with:

```
go build -tags lego_httpreq example.com/your/project
```

With this build constraint, your binary will only pull in the httpreq provider and not the much
larger providers such as GCP and AWS. In this particular case the binary is 30MB smaller.

## FAQ

### What's the point?

I compile most things I run in production. It upsets me that I'm compiling in support for a
million and one DNS providers that I'll never use, and otherwise very simple tools get ballooned
by 30MB or so because they're pulling in really complicated libraries for other DNS providers.

Obviously if you're building a general purpose release to be used in multiple environments
then letting the user select which one is probably the best way to go, but using legotapas
allows for advanced users to customise the build to their need while still maintaining the
general compatibility.

### Can I enable more than one provider?

No. I don't think this is possible with the way I'm using build constraints (without an
exponential explosion covering every possible combinations). I'm open to suggestions for
better ways to do this - please raise an issue or submit a PR!

### How to update when new providers are added

The application under `./cmd/generate` finds all available packages and executes templates
for each detected provider. Simply update the lego dependency with `go get` and then run
the generate command to refresh the templates.
