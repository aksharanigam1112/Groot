package routes

import (
	"github.com/aksharanigam1112/Groot/controllers/employee"
	"github.com/gin-gonic/gin"
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
