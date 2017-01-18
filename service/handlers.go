package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// getAllCatalogItemsHandler returns a fake list of catalog items
func getAllCatalogItemsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		catalog := make([]catalogItem, 2)
		catalog[0] = fakeItem("ABC123")
		catalog[1] = fakeItem("STAPLER99")
		formatter.JSON(w, http.StatusOK, catalog)
	}
}

// getCatalogItemDetailsHander returns a fake catalog item. The key here is that
// we're using a backing service to get fulfilment status for the individual item
func getCatalogItemDetailsHander(formatter *render.Render, serviceClient fulfillmentClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// returns the var variables in the request into vars
		vars := mux.Vars(r)
		// Simulates that we got a request to /catalog/{sku}
		sku := vars["sku"]
		fmt.Printf("Sku: %s \n", sku)
		// Fakes api call to fulfillment service
		status, err := serviceClient.getFulfillmentStatus(sku)
		// No errors
		if err == nil {
			// Formats to JSON and sends the response to the client
			formatter.JSON(w, http.StatusOK, catalogItem{
				ProductID:       1,
				SKU:             sku,
				Description:     "this is a fake product",
				Price:           1599, // 15.99
				ShipsWithin:     status.ShipsWithin,
				QuantityInStock: status.QuantityInStock,
			})
		} else {
			formatter.JSON(w, http.StatusInternalServerError, fmt.Sprintf("Fulfilment Client Error: %s", err.Error()))
		}
	}

}

func rootHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formatter.Text(w, http.StatusOK, "Catalog Service, see http://github.com/cloudnativego/backing-catalog for API.")
	}
}

func fakeItem(sku string) (item catalogItem) {
	item.SKU = sku
	item.Description = "This is a fake product"
	item.Price = 1599
	item.QuantityInStock = 75
	item.ShipsWithin = 14
	return item
}
