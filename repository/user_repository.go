package repository

import (
	"encoding/json"
	"go-api/model"
	"go-api/utils"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllProduct() ([]model.ProductCategoryWithProducts, error) {

	var rawData []model.ProductCategoryWithProducts
	err := r.DB.Select(&rawData, utils.GetAllProductActive)
	if err != nil {
		return nil, err
	}

	for i, row := range rawData {
		var products []model.Product
		if err := json.Unmarshal(row.ProductRaw, &products); err != nil {
			return nil, err
		}
		rawData[i].Product = products
	}
	return rawData, nil

}
