package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type PaymentClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewPaymentClient(baseURL string) *PaymentClient {
	return &PaymentClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 2 * time.Second,
		},
	}
}

func (c *PaymentClient) CheckPayment(orderID string, amount int64) (string, error) {
	requestBody, _ := json.Marshal(map[string]interface{}{
		"order_id": orderID,
		"amount":   amount,
	})

	resp, err := c.httpClient.Post(c.baseURL+"/payments", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "Failed", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "Failed", errors.New("payment service returned an error")
	}

	var result struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "Failed", err
	}

	return result.Status, nil
}
