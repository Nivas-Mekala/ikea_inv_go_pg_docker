package models

type Product struct {
	Product_Name string `gorm:"product_name"`
	Article_Id   string `gorm:"art_id"`
	Amount_Of    string `gorm:"amount_of"`
}

type ProductRequest struct {
	Products []struct {
		Name             string `json:"name"`
		Contain_Articles []struct {
			Article_Id string `json:"art_id"`
			Amount_Of  string `json:"amount_of"`
		} `json:"contain_articles"`
	} `json:"products"`
}
