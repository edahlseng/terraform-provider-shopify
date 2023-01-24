package shopifygo

import (
	"github.com/dghubble/sling"
)

type Client struct {
	sling    *sling.Sling
	Webhooks *WebhookService
}

func NewClient(shopifyDomain string, shopifyAccessToken string, shopifyApiVersion string) *Client {
	base := sling.New().Base("https://"+shopifyDomain).Set("Content-Type", "application/json").Set("X-Shopify-Access-Token", shopifyAccessToken)

	return &Client{
		sling:    base,
		Webhooks: newWebhookService(base.New(), shopifyApiVersion),
	}
}
