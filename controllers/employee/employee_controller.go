package employee

import (
	"net/http"
	"strconv"

	"github.com/aksharanigam1112/Groot/models"
	"github.com/aksharanigam1112/Groot/services"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	srvc services.Employee
}

func NewEmployeeController(srvc services.Employee) *EmployeeController {
	return &EmployeeController{srvc: srvc}
}

func (ctrl *EmployeeController) GetEmployeeDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	employeeId, _ := strconv.Atoi(id)

	result := ctrl.srvc.GetEmployeeById(employeeId)

	if result == nil || result.Id == 0 {
		ctx.IndentedJSON(http.StatusNotFound, models.Response{Message: "Data not found"})
		return
	}

	resp := models.Response{Employee: result}
	ctx.IndentedJSON(http.StatusOK, resp)
}
