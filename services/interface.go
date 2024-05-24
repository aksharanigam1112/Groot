package services

import "Groot/models"

type Employee interface {
	GetEmployeeById(int) *models.Employee
}
