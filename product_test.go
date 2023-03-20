package models

import (
	"bytes"
	"encoding/json"
	"github.com/Radser2001/products_api/routes"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	// create a mock product
	product := models.Product{Name: "Test Product", Price: "9.99"}

	// send a POST request to create the product
	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(product))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	// check the response status code
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200 but got %d", resp.StatusCode)
	}

	// check the response body
	body, _ := ioutil.ReadAll(resp.Body)
	var createdProduct routes.Product
	err := json.Unmarshal(body, &createdProduct)
	if err != nil {
		t.Errorf("Error unmarshaling response body: %s", err.Error())
	}
	if createdProduct.Name != product.Name {
		t.Errorf("Expected product name %s but got %s", product.Name, createdProduct.Name)
	}
	if createdProduct.Price != product.Price {
		t.Errorf("Expected product price %s but got %s", product.Price, createdProduct.Price)
	}
}

func TestGetProducts(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	// send a GET request to retrieve all products
	req, _ := http.NewRequest("GET", "/api/products", nil)
	resp, _ := app.Test(req)

	// check the response status code
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200 but got %d", resp.StatusCode)
	}

	// check the response body
	body, _ := ioutil.ReadAll(resp.Body)
	var products []routes.Product
	err := json.Unmarshal(body, &products)
	if err != nil {
		t.Errorf("Error unmarshaling response body: %s", err.Error())
	}
	if len(products) == 0 {
		t.Errorf("Expected at least one product but got none")
	}
}

func TestCreateProduct(t *testing.T) {
	app := fiber.New()
	routes.SetupRoutes(app)

	// create a mock product
	product := models.Product{Name: "Test Product", Price: "9.99"}

	// send a POST request to create the product
	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(product))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	// check the response status code
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200 but got %d", resp.StatusCode)
	}

	// check the response body
	body, _ := ioutil.ReadAll(resp.Body)
	var createdProduct routes.Product
	err := json.Unmarshal(body, &createdProduct)
	if err != nil {
		t.Errorf("Error unmarshaling response body: %s", err.Error())
	}
	if createdProduct.Name != product.Name {
		t.Errorf("Expected product name %s but got %s", product.Name, createdProduct.Name)
	}
	if createdProduct.Price != product.Price {
		t.Errorf("Expected product price %s but got %s", product.Price, createdProduct.Price)
	}
}
