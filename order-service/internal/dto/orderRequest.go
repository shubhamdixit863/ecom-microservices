package dto

type OrderRequest struct {
	ProductName        string `json:"product_name"`
	ProductID          int64  `json:"product_id"`
	ProductDescription string `json:"product_description"`
	ProductQuantity    int    `json:"product_quantity"`
	Address            string `json:"address"`
}
