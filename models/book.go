package models

type Book struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Price   string `json:"price"`
	Sales   string `json:"sales"`
	Stock   string `json:"stock"`
	ImgPath string `json:"img_path" gorm:"Column:img_path"`
}
