package types

import "time"

type ProductCategory struct {
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Product struct {
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	SKU         string    `json:"sku" db:"sku"`
	CategoryID  int       `json:"category_id" db:"category_id"`
	Price       int       `json:"price" db:"price"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type ProductInventory struct {
	ProductId int       `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ProductSummarized struct {
	Name                string    `json:"name" db:"name"`
	Description         string    `json:"description" db:"description"`
	SKU                 string    `json:"sku" db:"sku"`
	Price               int       `json:"price" db:"price"`
	CategoryName        string    `json:"category_name" db:"category_name"`
	CategoryDescription string    `json:"category_description" db:"category_description"`
	Quantity            string    `json:"quantity" db:"quantity"`
	Status              string    `json:"status" db:"status"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}
