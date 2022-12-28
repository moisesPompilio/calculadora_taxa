package database

import "database/sql"


type OrderRepository interface {
	Db *sql.DB

}
	func NewOrderRepository (db *sql.DB) *OrderRepository {
		return &OrderRepository {
            Db: db
        }
	}