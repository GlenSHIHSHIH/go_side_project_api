package forestage

type BaseForestageConfigDTO struct {
	BaseConfigDTO *BaseConfigDTO `form:"baseConfig" json:"baseConfig"`
}

type BaseConfigDTO struct {
	ImgUrl string `form:"imgUrl" json:"imgUrl"`
}
