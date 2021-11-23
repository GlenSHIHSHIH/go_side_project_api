package model

import "time"

const (
	UPDATE_SQL_SAFE_CLOSE    = `SET SQL_SAFE_UPDATES=0`
	UPDATE_SQL_SAFE_OPEN     = `SET SQL_SAFE_UPDATES=1`
	TRUNACTE_PRODUCTION_TEMP = `truncate table production_temps`

	UPDATE_PRODUCTION = `update productions as pd  INNER join production_temps as temp on pd.product_id = temp.product_id
							set
							pd.name=temp.name,
							pd.description=temp.description,
							pd.options=temp.options,
							pd.categories=temp.categories,
							pd.image=temp.image,
							pd.images=temp.images,
							pd.url=temp.url,
							pd.price=temp.price,
							pd.price_min=temp.price_min,
							pd.update_time=temp.update_time,
							pd.update_user_id=temp.update_user_id,
							pd.attribute=temp.attribute,
							pd.liked_count=temp.liked_count,
							pd.historical_sold=temp.historical_sold,
							pd.stock=temp.stock
							`

	INSERT_PRODUCTION = `insert into productions(product_id,name,description,options,categories,image,images,url,price,price_min,
						 create_time,update_time,create_user_id,attribute,liked_count,historical_sold,stock)
							select
								temp.product_id,temp.name,temp.description,temp.options,temp.categories,temp.image,temp.images,
								temp.url,temp.price,temp.price_min,temp.create_time,temp.update_time,temp.create_user_id,
								temp.attribute,temp.liked_count,temp.historical_sold,temp.stock
							from production_temps as temp left join productions as pd on pd.product_id = temp.product_id
							where  pd.id is null`
)

type ProductionTemp struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	ProductId      uint32    `gorm:"comment:產品ID(蝦皮ID);unique" json:"productId"`                                      //產品ID(蝦皮ID)
	Name           string    `gorm:"comment:產品名稱;type:varchar(100)" json:"name"`                                      //產品名稱
	Description    string    `gorm:"comment:敘述;type:mediumtext" json:"description"`                                   //敘述
	Options        string    `gorm:"comment:其他選項;type:text" json:"options"`                                           //其他選項
	Categories     string    `gorm:"comment:分類;type:varchar(100)" json:"categories"`                                  //分類
	Image          string    `gorm:"comment:主圖片;type:text" json:"image"`                                              //主圖片
	Images         string    `gorm:"comment:圖片(逗號區隔);type:text" json:"images"`                                        //圖片(逗號區隔)
	Url            string    `gorm:"comment:蝦皮連結;type:text" json:"url"`                                               //蝦皮連結
	Price          int64     `gorm:"comment:產品價格;type:int" json:"price"`                                              //產品價格
	PriceMin       int64     `gorm:"comment:產品價格低標;type:int" json:"priceMin"`                                         //產品價格低標
	Attribute      string    `gorm:"comment:其他商品資訊(json array);type:text" json:"attribute"`                           //其他商品資訊
	LikedCount     int       `gorm:"comment:喜歡的人數;type:int" json:"likedCount"`                                        //喜歡的人數
	HistoricalSold int       `gorm:"comment:銷售數量;type:int" json:"historicalSold"`                                     //銷售數量
	Stock          int       `gorm:"comment:商品庫存;type:int" json:"stock"`                                              //商品庫存
	CreateTime     time.Time `gorm:"comment:新增時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"` //新增時間
	UpdateTime     time.Time `gorm:"comment:更新時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"` //更新時間
	CreateUserId   int       `gorm:"comment:新增人員Id;type:int; json:"createUserId"`                                     //新增人員
	UpdateUserId   int       `gorm:"comment:修改人員Id;type:int; json:"updateUserId"`                                     //修改人員
}
