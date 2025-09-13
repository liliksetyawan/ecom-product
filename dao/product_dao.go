package dao

import (
	"database/sql"
	"ecom-product/dto/in"
	"ecom-product/repository"
	"log"
)

func CreateProduct(db *sql.DB, product *repository.ProductModel) (id int64, err error) {
	query := `
		INSERT INTO products (shop_id, code, name, description, price)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err = db.QueryRow(query,
		product.ShopID,
		product.Code,
		product.Name,
		product.Description,
		product.Price,
	).Scan(&id)

	if err != nil {
		log.Println("Error executing CreateProduct query:", err)
		return
	}
	return
}

// GetListProducts mengambil daftar produk dengan pagination dan search
func GetListProducts(db *sql.DB, p in.Pagination) ([]repository.ProductModel, error) {
	products := []repository.ProductModel{}

	query := `
		SELECT id, shop_id, code, name, description, price, created_at, updated_at
		FROM products
		WHERE ($1 = '' OR name ILIKE '%' || $1 || '%' OR code ILIKE '%' || $1 || '%')
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := db.Query(query, p.Search, p.Limit, (p.Offset-1)*p.Limit)
	if err != nil {
		log.Println("Error executing GetListProducts query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product repository.ProductModel
		if err := rows.Scan(
			&product.ID,
			&product.ShopID,
			&product.Code,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		); err != nil {
			log.Println("Error scanning product row:", err)
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductByID mengambil produk berdasarkan ID
func GetProductByID(db *sql.DB, id int64) (*repository.ProductModel, error) {
	query := `
		SELECT id, shop_id, code, name, description, price, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	var product repository.ProductModel
	err := db.QueryRow(query, id).Scan(
		&product.ID,
		&product.ShopID,
		&product.Code,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // tidak ada product
		}
		return nil, err
	}

	return &product, nil
}
