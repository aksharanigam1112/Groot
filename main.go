package main

import (
	"log"

	employeeController "github.com/aksharanigam1112/Groot/controllers/employee"
	"github.com/aksharanigam1112/Groot/database"
	employeeDao "github.com/aksharanigam1112/Groot/datasource/employee"
	"github.com/aksharanigam1112/Groot/routes"
	employeeSrvc "github.com/aksharanigam1112/Groot/services/employee"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// DB Client Initialization
	dbClient := database.NewMysqlClient()

	// Datasource Initialization
	repo := employeeDao.NewEmployeeRepository(dbClient)

	// Service Initialization
	srvc := employeeSrvc.NewEmployeeService(repo)

	// Controller Initialization
	ctrl := employeeController.NewEmployeeController(srvc)

	routes.AttachRoutes(router, ctrl)
	err := router.Run()
	if err != nil {
		log.Fatal("Error while running server")
	}
}
