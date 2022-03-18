package backstagedto

import (
	"componentmod/internal/dto"
	"componentmod/internal/dto/forestagedto"
	"time"
)

type CarouselListDTO struct {
	Carousel []*CarouselData           `json:"carousel"`
	PageData *dto.PageForMultSearchDTO `json:"pageData"`
}

type CarouselIdDTO struct {
	Carousel *CarouselData                   `json:"carousel"`
	Picture  []*forestagedto.PictureListData `json:"picture"`
}

type CarouselData struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`         //圖片名稱
	Weight       int       `json:"weight"`       //權重
	Status       bool      `json:"status"`       //狀態(開關)
	StartTime    time.Time `json:"startTime"`    //開始時間
	EndTime      time.Time `json:"endTime"`      //結束時間
	CreateTime   time.Time `json:"createTime"`   //新增時間
	UpdateTime   time.Time `json:"updateTime"`   //更新時間
	CreateUserId int       `json:"createUserId"` //新增人員
	UpdateUserId int       `json:"updateUserId"` //修改人員
}

type CarouselCreateOrEditDTO struct {
	Id        int                             `json:"id"`
	Name      string                          `json:"name"`      //圖片名稱
	Weight    int                             `json:"weight"`    //權重
	Status    bool                            `json:"status"`    //狀態(開關)
	StartTime time.Time                       `json:"startTime"` //開始時間
	EndTime   time.Time                       `json:"endTime"`   //結束時間
	Picture   []*forestagedto.PictureListData `json:"picture"`
}
