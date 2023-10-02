package repository

import (
	"database/sql"

	"github.com/devfullcycle/cqrs/internal/command/domain"
)

type OrderRepositoryMysql struct {
	DB *sql.DB
}

func (r *OrderRepositoryMysql) Save(order *domain.Order) {
	r.DB.Exec("INSERT INTO orders (id, total) VALUES (?, ?)", order.ID, order.Total)
	for _, orderItem := range order.OrderItems {
		r.DB.Exec("INSERT INTO order_items (id, order_id, product_id, quantity, price) VALUES (?, ?, ?, ?, ?)", orderItem.ID, order.ID, orderItem.Product.ID, orderItem.Quantity, orderItem.Price)
	}
}
