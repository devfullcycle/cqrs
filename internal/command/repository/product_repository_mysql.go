package repository

import (
	"database/sql"

	"github.com/devfullcycle/cqrs/internal/command/domain"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func (r *ProductRepositoryMysql) FindByID(id string) *domain.Product {
	var product domain.Product
	r.DB.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price)
	return &product
}

func (r *ProductRepositoryMysql) Save(product *domain.Product) {
	r.DB.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", product.ID, product.Name, product.Price)
}
