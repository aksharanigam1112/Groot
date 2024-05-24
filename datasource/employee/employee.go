package employee

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aksharanigam1112/Groot/models"
)

type EmployeeRepository struct {
	dbClient *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{
		dbClient: db,
	}
}

func (r *EmployeeRepository) GetById(id int) *models.Employee {

	rows, err := r.dbClient.Query(fmt.Sprintf("SELECT * FROM employee WHERE id = %d", id))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var data models.Employee
	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Name, &data.Role)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &data
}
