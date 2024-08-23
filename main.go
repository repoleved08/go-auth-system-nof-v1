package main

import (
	"go-auth-system/config"
	"go-auth-system/controllers"
	"log"
	"net/http"
)


func main() {
	// Load env
	config.LoadEnv()
	// Initialize DB
	config.InitDB()
	// Setup routes
	router := http.NewServeMux()
	router.HandleFunc("/products/create", controllers.AddProduct)
	router.HandleFunc("/products/all", controllers.GetProducts)
	router.HandleFunc("/products/one", controllers.GetProductById)
	router.HandleFunc("/products/update", controllers.UpdateProduct)
	router.HandleFunc("/products/delete", controllers.DeleteProduct)
	// protected routes
	// start server
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
