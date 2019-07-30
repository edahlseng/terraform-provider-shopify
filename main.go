package main

import (
	"github.com/edahlseng/terraform-provider-shopify/shopify"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: shopify.Provider})
}
