package main

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	_ "github.com/go-acme/lego/v4/providers/dns"
	"golang.org/x/tools/go/packages"
)

const dnsPackage = "github.com/go-acme/lego/v4/providers/dns"

//go:embed *.go.gotpl
var templates embed.FS

func main() {
	providerTemplate, err := template.ParseFS(templates, "provider.go.gotpl")
	if err != nil {
		panic(err)
	}

	p := providers()
	for i := range p {
		func() {
			f, err := os.Create(fmt.Sprintf("plate_%s.go", p[i]))
			if err != nil {
				panic(err)
			}
			defer f.Close()
			if err := providerTemplate.Execute(f, p[i]); err != nil {
				panic(err)
			}
		}()
	}

	defaultTemplate, err := template.New("default.go.gotpl").Funcs(template.FuncMap{"join": strings.Join}).ParseFS(templates, "default.go.gotpl")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("default.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := defaultTemplate.Execute(f, p); err != nil {
		panic(err)
	}
}

func providers() []string {
	ps, err := packages.Load(&packages.Config{Mode: packages.NeedImports}, dnsPackage)
	if err != nil {
		panic(err)
	}

	var res []string
	for i := range ps[0].Imports {
		imp := ps[0].Imports[i].ID
		if strings.HasPrefix(imp, dnsPackage) {
			res = append(res, path.Base(imp))
		}
	}

	sort.Strings(res)
	return res
}
