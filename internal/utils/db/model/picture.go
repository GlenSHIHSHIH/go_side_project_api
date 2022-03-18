package model

type Picture struct {
	Id       int        `gorm:"primaryKey" json:"id"`
	Name     string     `gorm:"comment:圖片名稱;type:varchar(100)" json:"name"`                //圖片名稱
	Alt      string     `gorm:"comment:圖片替代;type:text" json:"alt"`                         //圖片替代
	Url      string     `gorm:"comment:蝦皮連結;type:text" json:"url"`                         //連結
	Weight   int        `gorm:"comment:權重(優先順序 重=高);type:int;default:0" json:"weight"`     //權重
	Status   bool       `gorm:"comment:開關 (true=開啟);type:bool;default:true" json:"status"` //狀態(開關)
	Carousel []Carousel `gorm:"many2many:carousel_picture;"`                               //圖片(多對多)
}
