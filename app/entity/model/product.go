package model

import "github.com/golang-module/carbon/v2"

type ProductModels struct {
	CategoryId         int64         `json:"category_id"`
	ProductName        string        `json:"product_name"`
	ProductDescription string        `json:"product_description"`
	ProductPrice       int64         `json:"product_price"`
	ProductImage       string        `json:"product_image"`
	CreatedAt          carbon.Carbon `json:"created_at"`
	UpdatedAt          carbon.Carbon `json:"updated_at"`
}
