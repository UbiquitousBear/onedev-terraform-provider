package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/ubiquitousbear/onedev-terraform-provider/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
