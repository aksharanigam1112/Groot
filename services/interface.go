package services

import "github.com/aksharanigam1112/Groot/models"

type Employee interface {
	GetEmployeeById(int) *models.Employee
}
