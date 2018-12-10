package shopify

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHOPIFY_DOMAIN", nil),
				Description: "Domain of the Shopify store",
			},
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHOPIFY_ACCESS_TOKEN", nil),
				Description: "Shopify access token",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"shopify_webhook": resourceShopifyWebhook(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ShopifyDomain:      d.Get("domain").(string),
		ShopifyAccessToken: d.Get("access_token").(string),
	}

	client := config.NewClient()

	return client, nil
}
