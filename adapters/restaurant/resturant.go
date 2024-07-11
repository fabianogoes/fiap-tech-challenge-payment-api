package restaurant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fabianogoes/fiap-payment/domain/entities"
)

type ClientAdapter struct {
	config *entities.Config
}

func NewClientAdapter(config *entities.Config) ClientAdapter {
	return ClientAdapter{config: config}
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
	url := fmt.Sprintf("%s/orders/%d/payment/webhook", p.config.RestaurantApiUrl, orderID)
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
	log.Println(sb)

	return nil
}
