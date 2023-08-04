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
	return &DB{Conn: Conn}
}

// Implement GetAll customer for repository.
func (db *DB) GetAll() (*[]models.Customers, error) {
	rows, errExec := db.Conn.Query(`SELECT id, name, phone, email, age FROM customers`)
	if errExec != nil {
		return nil, errExec
	}

	var result []models.Customers

	// Loop through rows customers
	for rows.Next() {
		var each models.Customers
		// scan will input each field customer
		errScan := rows.Scan(&each.Id, &each.Name, &each.Phone, &each.Email, &each.Age)

		if errScan != nil {
			return nil, errScan
		}
		result = append(result, each)
	}

	return &result, nil
}

// Implement Create customer for repository.
func (db *DB) Create(c *models.RequestInsertCustomer) error {
	stmt, err := db.Conn.Prepare(`INSERT INTO customers (name, phone, email, age) VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return err
	}
	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age)
	if err != nil {
		return errExec
	}
	return nil
}

// Implement Update customer for repository.
func (db *DB) Update(c *models.RequestUpdateCustomer) error {
	stmt, err := db.Conn.Prepare(`UPDATE customers SET name = $1, phone = $2, email = $3, age = $4 WHERE id = $5`)
	if err != nil {
		return err
	}
	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age, c.Id)
	if err != nil {
		return errExec
	}
	return nil
}

// Implement Delete customer for repository.
func (db *DB) Delete(Id uint64) error {
	_, errExec := db.Conn.Exec(`DELETE FROM customers WHERE id = $1`, Id)
	if errExec != nil {
		return errExec
	}
	return nil
}