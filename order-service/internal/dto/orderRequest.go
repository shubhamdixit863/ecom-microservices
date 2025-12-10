package dto

type OrderRequest struct {
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductQuantity    int    `json:"product_quantity"`
	Address            string `json:"address"`
}
