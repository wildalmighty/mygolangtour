package product

import (
	"encoding/json"
	"fmt"
	"github.com/wildalmighty/mygolangtour/rest/middleware"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const productsBasePath = "products"

func SetupRoutes(apiBasePath string) {
	handleProduct := http.HandlerFunc(productHandler)
	handleProducts := http.HandlerFunc(productsHandler)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsBasePath), middleware.DurationHandler(handleProducts))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsBasePath), middleware.DurationHandler(handleProduct))
}

func productHandler(writer http.ResponseWriter, request *http.Request) {
	urlPathSegment := strings.Split(request.URL.Path, "products/")
	productID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	product := getProduct(productID)
	if product == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	switch request.Method {
	case http.MethodGet:
		productJson, err := json.Marshal(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(productJson)
	case http.MethodPut:
		//update product in the list
		newProduct, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		var updatedProduct Product
		err = json.Unmarshal(newProduct, &updatedProduct)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedProduct.ProductID != productID {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		addOrUpdateProduct(updatedProduct)
		writer.WriteHeader(http.StatusOK)
		return
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productList := getProductList()
		productsJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)
	case http.MethodPost:
		//add a new product to the list
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdateProduct(newProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
}