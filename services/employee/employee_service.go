package employee

import (
	"github.com/aksharanigam1112/Groot/datasource"
	"github.com/aksharanigam1112/Groot/models"
)

type EmployeeService struct {
	repo datasource.Employee
}

func NewEmployeeService(employee datasource.Employee) *EmployeeService {
	return &EmployeeService{repo: employee}
}

func (srvc *EmployeeService) GetEmployeeById(id int) *models.Employee {
	return srvc.repo.GetById(id)
}
