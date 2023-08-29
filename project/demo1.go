package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func Demo1() {
	//GetAllProducts()
	//AddProducts()
	products, _ := GetAllProducts()
	fmt.Println("Urunler := ", products)
	for _, v := range products {
		fmt.Println(v.ProductName)
	}
	var defaultProd = Product{}
	addedProduct, _ := AddProducts()
	if addedProduct == defaultProd {
		fmt.Printf("urun eklenemedi")
	} else {
		fmt.Println("Eklendi", addedProduct)
	}

}
func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil

}
func AddProducts() (Product, error) {
	product := Product{Id: 11, ProductName: "Mekanik Klavye", CategoryId: 1, UnitPrice: 1000.00}
	jsonProduct, err := json.Marshal(product)
	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		return Product{}, err
	}
	defer response.Body.Close()
	bodybytes, _ := ioutil.ReadAll(response.Body)
	var productResponse Product
	json.Unmarshal(bodybytes, &productResponse)
	fmt.Println("kaydedildi", productResponse)
	return productResponse, nil
}
