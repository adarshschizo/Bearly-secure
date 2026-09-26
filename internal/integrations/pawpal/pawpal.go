package pawpal

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
)

type WebhookOutcome string

const (
	WebhookUnauthorized WebhookOutcome = "unauthorized"
	WebhookMalformed    WebhookOutcome = "malformed"
	WebhookApproved     WebhookOutcome = "approved"
)

type WebhookVerification struct {
	Outcome WebhookOutcome
	OrderID int64
}

func CreateCheckoutURL(orderID int64) string {
	return fmt.Sprintf("https://pawpal.example/checkout?orderId=%d", orderID)
}

func VerifyWebhook(providedKey, expectedKey string, payload []byte) WebhookVerification {
	if len(providedKey) != len(expectedKey) ||
		subtle.ConstantTimeCompare([]byte(providedKey), []byte(expectedKey)) != 1 {
		return WebhookVerification{
			Outcome: WebhookUnauthorized,
		}
	}

	var webhook struct {
		OrderID int64  `json:"orderId"`
		Status  string `json:"status"`
	}

	if err := json.Unmarshal(payload, &webhook); err != nil {
		return WebhookVerification{
			Outcome: WebhookMalformed,
		}
	}

	if webhook.OrderID <= 0 || webhook.Status != "approved" {
		return WebhookVerification{
			Outcome: WebhookMalformed,
		}
	}

	return WebhookVerification{
		Outcome: WebhookApproved,
		OrderID: webhook.OrderID,
	}
}
