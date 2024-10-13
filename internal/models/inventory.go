package models

type Inventory struct {
	Article_Id string `json:"art_id" gorm:"article_id; primaryKey"`
	Name       string `json:"name" gorm:"name"`
	Stock      string `json:"stock" gorm:"stock"`
}

type InventoryRequest struct {
	Inventory []struct {
		Article_Id string `json:"art_id"`
		Name       string `json:"name"`
		Stock      string `json:"stock"`
	} `json:"inventory"`
}
