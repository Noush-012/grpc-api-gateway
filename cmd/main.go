package main

import (
	"fmt"
	"log"

	"github.com/Noush-012/grpc-microservice/pkg/auth"
	"github.com/Noush-012/grpc-microservice/pkg/config"
	"github.com/Noush-012/grpc-microservice/pkg/order"
	"github.com/Noush-012/grpc-microservice/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	fmt.Println("-------", c, c.ProductSvcUrl)

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
