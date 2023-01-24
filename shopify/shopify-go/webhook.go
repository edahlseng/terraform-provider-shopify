package shopifygo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dghubble/sling"
)

type WebhookService struct {
	sling *sling.Sling
}

func newWebhookService(sling *sling.Sling, shopifyApiVersion string) *WebhookService {
	return &WebhookService{
		sling: sling.New().Path("admin/api/" + shopifyApiVersion + "/"),
	}
}

// -----------------------------------------------------------------------------
// Input
// -----------------------------------------------------------------------------

type WebhookInput struct {
	Topic   string `json:"topic"`
	Address string `json:"address"`
	Format  string `json:"format"`
}

type WebhookInputBody struct {
	Webhook WebhookInput `json:"webhook"`
}

// -----------------------------------------------------------------------------
// Responses
// -----------------------------------------------------------------------------

type Webhook struct {
	Id         int    `json:"id"`
	Topic      string `json:"topic"`
	Address    string `json:"address"`
	Format     string `json:"format"`
	Created_at string `json:"created_at"`
	Updated_at string `json."updated_at"`
}

type WebhookResponse struct {
	Webhook Webhook `json:"webhook"`
}

type WebhookDeleteResponse struct{}

// -----------------------------------------------------------------------------
// Errors
// -----------------------------------------------------------------------------

type WebhookErrorMessage struct {
	WebhookMessage *string   `json:"webhook,omitempty"`
	TopicMessage   *[]string `json:"topic,omitempty"`
	AddressMessage *[]string `json:"address,omitempty"`
	FormatMessage  *[]string `json:"format,omitempty"`
}

type WebhookError struct {
	Errors WebhookErrorMessage `json:"errors"`
}

func (err WebhookError) Error() string {
	if err == (WebhookError{}) {
		return ""
	}

	webhookMessage := ""
	if err.Errors.WebhookMessage != nil {
		webhookMessage = " " + *err.Errors.WebhookMessage
	}

	topicMessage := ""
	if err.Errors.TopicMessage != nil {
		topicMessage = fmt.Sprintf(" topic: %v", *err.Errors.TopicMessage)
	}

	addressMessage := ""
	if err.Errors.AddressMessage != nil {
		addressMessage = fmt.Sprintf(" address: %v", *err.Errors.AddressMessage)
	}

	formatMessage := ""
	if err.Errors.FormatMessage != nil {
		formatMessage = fmt.Sprintf(" format: %v", *err.Errors.FormatMessage)
	}

	return fmt.Sprintf("Shopify:%s%s%s%s", webhookMessage, topicMessage, addressMessage, formatMessage)
}

// -----------------------------------------------------------------------------
// CRUD
// -----------------------------------------------------------------------------

func (service *WebhookService) Create(params *WebhookInput) (Webhook, *http.Response, error) {
	webhookResponse := new(WebhookResponse)
	genericError := new(json.RawMessage)
	resp, err := service.sling.New().Post("webhooks.json").BodyJSON(&WebhookInputBody{Webhook: *params}).Receive(webhookResponse, genericError)
	return webhookResponse.Webhook, resp, relevantError(resp, err, genericError, new(WebhookError))
}

func (service *WebhookService) Read(webhookId string) (Webhook, *http.Response, error) {
	webhookResponse := new(WebhookResponse)
	genericError := new(json.RawMessage)
	path := fmt.Sprintf("webhooks/%s.json", url.PathEscape(webhookId))
	resp, err := service.sling.New().Get(path).Receive(webhookResponse, genericError)
	return webhookResponse.Webhook, resp, relevantError(resp, err, genericError, new(WebhookError))
}

func (service *WebhookService) Update(webhookId string, params *WebhookInput) (Webhook, *http.Response, error) {
	webhookResponse := new(WebhookResponse)
	genericError := new(json.RawMessage)
	path := fmt.Sprintf("webhooks/%s.json", url.PathEscape(webhookId))
	resp, err := service.sling.New().Put(path).BodyJSON(&WebhookInputBody{Webhook: *params}).Receive(webhookResponse, genericError)
	return webhookResponse.Webhook, resp, relevantError(resp, err, genericError, new(WebhookError))
}

func (service *WebhookService) Delete(webhookId string) (*http.Response, error) {
	webhookDeleteResponse := new(WebhookDeleteResponse)
	genericError := new(json.RawMessage)
	path := fmt.Sprintf("webhooks/%s.json", url.PathEscape(webhookId))
	resp, err := service.sling.New().Delete(path).Receive(webhookDeleteResponse, genericError)

	return resp, relevantError(resp, err, genericError, new(WebhookError))
}
