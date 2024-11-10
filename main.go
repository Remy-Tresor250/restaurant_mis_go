package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"golang_restaurant_mis/database"
	"golang_restaurant_mis/routes"
	"golang_restaurant_mis/middleware"
	"go.mongodb.org/mongo-driver/mongo"
);

func main() {
	port := os.Getenv("PORT");

	if port == "" {
		port = "8000"
	}

	
}