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

func (p *ClientAdapter) Webhook(orderID uint, status string, method string) error {
	fmt.Printf("Payment webhook request for order %d status %s method %s \n", orderID, status, method)

	postBody, _ := json.Marshal(map[string]interface{}{
		"status":        status,
		"paymentMethod": method,
	})
	fmt.Printf("PUT body: %s\n", string(postBody))

	url := fmt.Sprintf("%s/orders/%d/payment/webhook", p.config.RestaurantApiUrl, orderID)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalf("An Error Occured to prepar request %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("calling restaurant webhook url %s \n", url)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An Error Occured to call restaurant webhook %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An Error Occured to read response body %v", err)
		return err
	}

	sb := string(body)
	log.Println(sb)

	return nil
}
