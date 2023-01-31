package main

import (
	"github.com/plurasight/webservice/database"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/plurasight/webservice/product"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
