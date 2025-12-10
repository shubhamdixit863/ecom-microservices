package models

import "time"

type Order struct {
	ID                 string    `json:"id" bson:"_id,omitempty"`
	ProductName        string    `json:"product_name" bson:"product_name"`
	ProductDescription string    `json:"product_description" bson:"product_description"`
	ProductQuantity    int       `json:"product_quantity" bson:"product_quantity"`
	Address            string    `json:"address" bson:"address"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
}
