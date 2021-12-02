package model

type Toko struct {
	ID            int    `json:"id"`
	Restaurant_id string `json:"restaurant_id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	PictureID     string `json:"pictureID"`
	City          string `json:"city"`
	Rating        string `json:"rating"`
}
