package datasource

import "github.com/aksharanigam1112/Groot/models"

type Employee interface {
	GetById(int) *models.Employee
}
