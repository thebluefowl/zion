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
	URL          string
	Headers      map[string]string
	SigningToken string
}

func (w *Webhook) Sign(payload []byte) string {
	key := []byte(w.SigningToken)
	mac := hmac.New(sha256.New, key)
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}

func (w *Webhook) Send(payload []byte, responseCallback func([]byte, error) error) error {
	if w.SigningToken != "" {
		w.Headers["X-Zion-Signature"] = fmt.Sprintf("v0=%s", w.Sign(payload))
	}
	request, _ := http.NewRequest("POST", w.URL, bytes.NewBuffer(payload))
	for k, v := range w.Headers {
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
