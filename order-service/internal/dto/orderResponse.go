package dto

type OrderResponse struct {
	ID                 string `json:"id"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductQuantity    int    `json:"product_quantity"`
	Address            string `json:"address"`
}
