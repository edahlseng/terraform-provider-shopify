package shopify

import (
	shopify "./shopify-go"
)

type Config struct {
	ShopifyDomain      string
	ShopifyAccessToken string
}

func (c *Config) NewClient() *shopify.Client {
	return shopify.NewClient(c.ShopifyDomain, c.ShopifyAccessToken)
}
