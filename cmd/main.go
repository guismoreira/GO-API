package main

import (
	"go-API/controller"
	"go-API/db"
	"go-API/repository"
	"go-API/usecase"

	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if(err != nil) {
		panic(err)
	}

	ProductRepository := repository.NewProductController(dbConnection)
	ProductUsecase := usecase.NewProductUseCase(ProductRepository)
	ProductControler := controller.NewProductController(ProductUsecase)

	server.GET("/products", ProductControler.GetProducts)
	server.POST("/product", ProductControler.CreateProduct)
	server.GET("/product/:productId", ProductControler.GetProductById)

	server.Run(":8000")
}