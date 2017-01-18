package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type fulfillmentClient interface {
	getFulfillmentStatus(sku string) (status fulfillmentStatus, err error)
}

type fulfillmentWebClient struct {
	rootURL string
}

func (client fulfillmentWebClient) getFulfillmentStatus(sku string) (status fulfillmentStatus, err error) {
	httpclient := &http.Client{}

	skuURL := fmt.Sprintf("%s%s", client.rootURL, sku)
	fmt.Printf("About to request SKU details from backing service: %s \n", skuURL)
	req, _ := http.NewRequest("GET", skuURL, nil)
	resp, err := httpclient.Do(req)

	if err != nil {
		fmt.Printf("Errored when sending request to server: %s \n", err)
		return
	}
	defer resp.Body.Close()

	payload, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(payload, &status)
	if err != nil {
		fmt.Printf("Failed to unmarshal server response")
		return
	}
	return status, err
}
