package customerrepository

import (
	"database/sql"

	"github.com/aldimaulana1605/bootcamp-api-hmsi/models"
	"github.com/aldimaulana1605/bootcamp-api-hmsi/modules/customers"
)

type DB struct {
	Conn *sql.DB
}

func NewCustomerRepository(Conn *sql.DB) customers.CustomerRepository {
	return &DB{Conn}
}
	
func (db *DB) GetAll() (*[]models.Customers, error) {
	rows, errExec := db.Conn.Query(`SELECT id, name, phone, email, age FROM customers`)
	if errExec != nil {
		return nil, errExec
	}
	var result []models.Customers
	for rows.Next() {
		var each models.Customers

		errScan := rows.Scan(&each.Id, &each.Name, &each.Phone, &each.Email, &each.Age)
		if errScan!= nil {
            return nil, errScan
        }
		result = append(result, each)
	}
	return &result, nil

	
}

func (db *DB) Create(c *models.RequestInsertCustomer) error {
	stmt, err := db.Conn.Prepare(`INSERT INTO customers (name, phone, email, age) VALUES ($1, $2, $3, $4)`)

	if err != nil {
		return err
	}

	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age)

	if errExec != nil {
		return errExec
	}

	return nil
}