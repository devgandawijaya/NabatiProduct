package model

import "encoding/json"

type Product struct {
	ID          int     `db:"id" json:"id"`
	CategoryID  int     `db:"category_id" json:"category_id"`
	Name        string  `db:"name" json:"name"`
	SKU         string  `db:"sku" json:"sku"`
	Description string  `db:"description" json:"description"`
	Price       float64 `db:"price" json:"price"`
	PromoPrice  float64 `db:"promo_price" json:"promo_price"`
	Stock       int     `db:"stock" json:"stock"`
	Weight      float64 `db:"weight" json:"weight"`
	ImageURL    string  `db:"image_url" json:"image_url"`
	IsActive    bool    `db:"is_active" json:"is_active"`
	CreatedAt   string  `db:"created_at" json:"created_at"`
	UpdatedAt   string  `db:"updated_at" json:"updated_at"`
}

type ProductCategoryWithProducts struct {
	ID         int             `db:"id_category" json:"id_category"`
	Name       string          `db:"name_category" json:"name_category"`
	Deskripsi  string          `db:"deskripsi" json:"deskripsi"`
	ProductRaw json.RawMessage `db:"product" json:"-"`
	Product    []Product       `json:"product"`
}
