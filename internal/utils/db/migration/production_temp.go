package migration

import "time"

type ProductionTemp struct {
	Id          int       `gorm:"primaryKey" json:"id"`
	ProductId   uint32    `gorm:"comment:產品ID(蝦皮ID);unique" json:"productId"`                                      //產品ID(蝦皮ID)
	Name        string    `gorm:"comment:產品名稱;type:varchar(100)" json:"name"`                                      //產品名稱
	Description string    `gorm:"comment:敘述;type:mediumtext" json:"descriptio"`                                    //敘述
	Options     string    `gorm:"comment:其他選項;type:text" json:"options"`                                           //其他選項
	Categories  string    `gorm:"comment:分類;type:varchar(100)" json:"categories"`                                  //分類
	Image       string    `gorm:"comment:主圖片;type:text" json:"image"`                                              //主圖片
	Images      string    `gorm:"comment:圖片(逗號區隔);type:text" json:"images"`                                        //圖片(逗號區隔)
	Url         string    `gorm:"comment:蝦皮連結;type:text" json:"url"`                                               //蝦皮連結
	CreateTime  time.Time `gorm:"comment:新增時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"` //新增時間
	UpdateTime  time.Time `gorm:"comment:更新時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"` //更新時間
}
