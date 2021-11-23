package model

import "time"

// `DROP PROCEDURE IF EXISTS get_prod_categories;
// DELIMITER //
// CREATE PROCEDURE get_prod_categories()
// BEGIN

// set @str_categories=  TRIM( BOTH  ',' FROM  (select  replace(GROUP_CONCAT(categories),'、',',') from productions));

// select  distinct substring_index(substring_index(@str_categories,',',help_topic_id+1),',',-1) as categories
// from mysql.help_topic
// where help_topic_id < (length(@str_categories)-length(replace(@str_categories,',',''))+1)
// order by categories asc;

// END
// //
// DELIMITER ;`

const (
	DROP_PROCEDURE_IF_EXISTS      = `DROP PROCEDURE IF EXISTS get_prod_categories`
	PROCEDURE_GET_PROD_CATEGORIES = `CREATE PROCEDURE get_prod_categories()
									BEGIN
									set @str_categories=  TRIM( BOTH  ',' FROM  (select  replace(GROUP_CONCAT(categories),'、',',') from productions));
									select  distinct substring_index(substring_index(@str_categories,',',help_topic_id+1),',',-1) as categories
									from mysql.help_topic
									where help_topic_id < (length(@str_categories)-length(replace(@str_categories,',',''))+1)
									order by categories asc;
									END`
	GET_PROD_CATEGORIES = `CALL get_prod_categories();`
)

type Production struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	ProductId      uint32    `gorm:"comment:產品ID(蝦皮ID);unique" json:"productId"`                                      //產品ID(蝦皮ID)
	Name           string    `gorm:"comment:產品名稱;type:varchar(100)" json:"name"`                                      //產品名稱
	Description    string    `gorm:"comment:敘述;type:mediumtext" json:"descriptio"`                                    //敘述
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
	Weight         int       `gorm:"comment:權重(優先順序 重=高);type:int;default:0" json:"weight"`                           //權重
	CreateTime     time.Time `gorm:"comment:新增時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"` //新增時間
	UpdateTime     time.Time `gorm:"comment:更新時間;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updateTime"` //更新時間
	CreateUserId   int       `gorm:"comment:新增人員Id;type:int; json:"createUserId"`                                     //新增人員
	UpdateUserId   int       `gorm:"comment:修改人員Id;type:int; json:"updateUserId"`                                     //修改人員
}
