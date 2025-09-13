package repository

import "database/sql"

type ProductModel struct {
	ID          sql.NullInt64
	ShopID      sql.NullInt64
	Code        sql.NullString
	Name        sql.NullString
	Description sql.NullString
	Price       sql.NullFloat64
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}
