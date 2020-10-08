package models

type Book struct {
	ID      int    `json:"id"`
	Title   string `json:"title" gorm:"Column:title; type: varchar(255)" form:"title"`
	Author  string `json:"author" gorm:"Column:author; type: varchar(255)" form:"author"`
	Price   string `json:"price" gorm:"Column:price; type: decimal(10,2)" form:"price"`
	Sales   string `json:"sales" gorm:"Column:sales; type: int(11)" form:"sales"`
	Stock   string `json:"stock" gorm:"Column:stock; type: int(11)" form:"stock"`
	ImgPath string `json:"img_path" gorm:"Column:img_path" form:"img_path"`
}

func (Book) TableName() string {
	return "books"
}
