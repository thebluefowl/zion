package dispatcher

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Webhook struct {
	config *WebhookConfig
}

type WebhookConfig struct {
	URL          string            `json:"url"`
	Headers      map[string]string `json:"headers"`
	SigningToken string            `json:"signing_token"`
}

func NewWebhook(config interface{}) *Webhook {
	return &Webhook{config: config.(*WebhookConfig)}
}

func (w *Webhook) Sign(payload []byte) string {
	key := []byte(w.config.SigningToken)
	mac := hmac.New(sha256.New, key)
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}

func (w *Webhook) Send(payload []byte, responseCallback func([]byte, error) error) error {
	if w.config.SigningToken != "" {
		w.config.Headers["X-Zion-Signature"] = fmt.Sprintf("v0=%s", w.Sign(payload))
	}
	request, _ := http.NewRequest("POST", w.config.URL, bytes.NewBuffer(payload))
	for k, v := range w.config.Headers {
		request.Header.Set(k, v)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err := responseCallback(body, err); err != nil {
		return err
	}
	return nil
}
