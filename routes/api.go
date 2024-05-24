package routes

import (
	"github.com/gin-gonic/gin"
	"Groot/controllers/employee"
)

const (
	HEALTH_CHECK = "health"
	GET_EMPLOYEE = "v1/employee/:id"
)

// AttachRoutes : attaches v2 apis
func AttachRoutes(r *gin.Engine, ctrl *employee.EmployeeController) {
	api := r.Group("/")
	{
		api.GET(GET_EMPLOYEE, ctrl.GetEmployeeDetails)
	}
}
