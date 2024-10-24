package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductController struct{}

var products = []Product{
	{ID: "1", Name: "Product 1"},
	{ID: "2", Name: "Product 2"},
}

// Função que lida com a requisição para listar os produtos
func (pc ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Função que lida com a requisição para obter um produto pelo ID
func (pc ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, product := range products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	// Instanciar o controlador
	productController := ProductController{}

	// Configurar as rotas
	http.HandleFunc("/products", productController.GetProducts)
	http.HandleFunc("/product", productController.GetProductByID)

	// Iniciar o servidor
	http.ListenAndServe(":8080", nil)
}
