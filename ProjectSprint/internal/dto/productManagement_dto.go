package dto

type ProductRequest struct {
	Name           string   `json:"name" validate:"required,min=5,max=60"`
	Price          float64  `json:"price" validate:"required,min=0"`
	ImageURL       string   `json:"imageUrl" validate:"required,url"`
	Stock          int      `json:"stock" validate:"required,min=0"`
	Condition      string   `json:"condition" validate:"required,oneof=new second"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable" validate:"required"`
}
