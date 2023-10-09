package models

type ProductMedia struct {
	ID        uint `json:"id"`
	ProductID int `json:"productId"`
	MediaID   int `json:"mediaId"`
}
