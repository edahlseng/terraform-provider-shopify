package shopify

import (
	shopify "github.com/edahlseng/terraform-provider-shopify/shopify/shopify-go"
)

type Config struct {
	ShopifyDomain      string
	ShopifyAccessToken string
	ShopifyApiVersion  string
}

func (c *Config) NewClient() *shopify.Client {
	return shopify.NewClient(c.ShopifyDomain, c.ShopifyAccessToken, c.ShopifyApiVersion)
}
