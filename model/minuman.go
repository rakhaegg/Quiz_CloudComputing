package model

type Minuman struct {
	Restaurant_id string `json:"restaurant_id"`
	Picture       string `json:"picture"`
	Description   string `json:"description"`
	Name          string `json:"name"`
	Harga         int    `json:"harga"`
}
