package datasource

import "Groot/models"

type Employee interface {
	GetById(int) *models.Employee
}
