package controllers

import (
	"encoding/json"
	"go-auth-system/config"
	"go-auth-system/models"
	"log"
	"net/http"
	"strconv"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
	}

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := `INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id`
	err = config.DB.QueryRow(query, product.Name, product.Description, product.Price).Scan(&product.ID)
	if err != nil {
		http.Error(w, "error when adding product", http.StatusInternalServerError)
		log.Printf("Error %v", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product added successfully"))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid request method", http.StatusBadRequest)
		return
	}
	rows, err := config.DB.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		http.Error(w, "error fetching products", http.StatusInternalServerError)
		log.Printf("error: %v", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid request method", http.StatusBadRequest)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "an error occurred", http.StatusInternalServerError)
		return
	}
	var product models.Product
	query := "SELECT id, name, description, price FROM products WHERE id=$1"
	err = config.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		http.Error(w, "Product Not Found!", http.StatusNotFound)
		log.Printf("Error %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := "UPDATE products SET name=$1, description=$2, price=$3 WHERE id=$4"
	_, err = config.DB.Exec(query, product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product updated successfully"))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "parameter id is missing", http.StatusBadRequest)
		return
	}
	query := "DELETE FROM products WHERE id=$1"
	_, err = config.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Product deletion failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
}
