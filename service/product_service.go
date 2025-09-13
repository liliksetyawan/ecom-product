package service

import (
	"database/sql"
	"ecom-product/dao"
	"ecom-product/dto/in"
	"ecom-product/dto/out"
	"ecom-product/repository"
	"ecom-product/server"
	"fmt"
	"log"
	"time"
)

func CreateProduct(req *in.ProductDTOIn) (*out.ProductDTOOut, error) {
	if req.ShopID == 0 || req.Code == "" || req.Name == "" || req.Price <= 0 {
		return nil, fmt.Errorf("shop_id, code, name, and valid price are required")
	}

	product := &repository.ProductModel{
		ShopID:      sql.NullInt64{Int64: req.ShopID, Valid: true},
		Code:        sql.NullString{String: req.Code, Valid: true},
		Name:        sql.NullString{String: req.Name, Valid: true},
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Price:       sql.NullFloat64{Float64: req.Price, Valid: true},
	}

	db := server.DBConn
	id, err := dao.CreateProduct(db, product)
	if err != nil {
		return nil, fmt.Errorf("error saving product: %v", err)
	}

	return &out.ProductDTOOut{
		ID:          id,
		ShopID:      req.ShopID,
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}, nil
}

func GetProductByID(id int64) (*out.ProductDTOOut, error) {
	db := server.DBConn
	productModel, err := dao.GetProductByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}

	return &out.ProductDTOOut{
		ID:          productModel.ID.Int64,
		ShopID:      productModel.ShopID.Int64,
		Code:        productModel.Code.String,
		Name:        productModel.Name.String,
		Description: productModel.Description.String,
		Price:       productModel.Price.Float64,
		CreatedAt:   productModel.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   productModel.UpdatedAt.Time.Format(time.RFC3339),
	}, nil
}

func GetProducts(req *in.GetListDTO) ([]*out.ProductDTOOut, error) {
	db := server.DBConn

	pagination := in.Pagination{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: req.Search,
	}

	products, err := dao.GetListProducts(db, pagination)
	if err != nil {
		log.Println("Error fetching products:", err)
		return nil, err
	}

	var productDTOs []*out.ProductDTOOut
	for _, productModel := range products {
		productDTO := &out.ProductDTOOut{
			ID:          productModel.ID.Int64,
			ShopID:      productModel.ShopID.Int64,
			Code:        productModel.Code.String,
			Name:        productModel.Name.String,
			Description: productModel.Description.String,
			Price:       productModel.Price.Float64,
			CreatedAt:   productModel.CreatedAt.Time.Format(time.RFC3339),
			UpdatedAt:   productModel.UpdatedAt.Time.Format(time.RFC3339),
		}
		productDTOs = append(productDTOs, productDTO)
	}

	return productDTOs, nil
}
