package models

type Condition string

const (
	New    Condition = "new"
	Second Condition = "second"
)

type Product struct {
	ProductID      string    `json:"productId"`
	Name           string    `json:"name"`
	Price          int       `json:"price"`
	ImageUrl       string    `json:"imageUrl"`
	Stock          int       `json:"stock"`
	Condition      Condition `json:"condition"`
	IsPurchaseable bool      `json:"isPurchaseable"`
	UserID         string    `json:"userId"`
}

type Tag struct {
	TagID string `json:"tagId"`
	Name  string `json:"name"`
}

type ItemTag struct {
	ProductID string `json:"productId"`
	TagID     string `json:"tagId"`
}
