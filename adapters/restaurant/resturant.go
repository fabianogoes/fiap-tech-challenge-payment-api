package restaurant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ClientAdapter struct {
}

func NewClientAdapter() *ClientAdapter {
	return &ClientAdapter{}
}

type PaymentWebhookRequest struct {
	Status        string `json:"status"`
	PaymentMethod string `json:"paymentMethod"`
	ErrorReason   string `json:"errorReason,omitempty"`
}

func (p *ClientAdapter) Webhook(orderID uint, status string) error {
	fmt.Printf("Payment webhook request for order %d status %s \n", orderID, status)

	postBody, _ := json.Marshal(map[string]interface{}{
		"status": status,
		"method": status,
	})
	fmt.Printf("Post body: %s\n", string(postBody))

	responseBody := bytes.NewBuffer(postBody)
	url := fmt.Sprintf("http://localhost:8080/orders/%d/payment/webhook", orderID)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)

	return nil
}
