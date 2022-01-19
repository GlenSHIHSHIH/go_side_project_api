package backstagectl

import "componentmod/internal/dto"

func GetPageMultSearchDefaultDTO() *dto.PageForMultSearchDTO {
	pageForMultSearchDTO := &dto.PageForMultSearchDTO{
		Page:       1,
		PageLimit:  20,
		Sort:       "asc",
		SortColumn: "id",
		Search:     make(map[string]string),
	}
	return pageForMultSearchDTO
}
