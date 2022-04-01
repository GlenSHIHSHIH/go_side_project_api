package backstage

import (
	"componentmod/internal/api/config"
	errorcode "componentmod/internal/api/errorcode"
	"componentmod/internal/dto"
	"componentmod/internal/dto/backstagedto"
	"componentmod/internal/dto/forestagedto"
	"componentmod/internal/services/api"
	"componentmod/internal/utils"
	"componentmod/internal/utils/db"
	"componentmod/internal/utils/db/model"
	"componentmod/internal/utils/file"
	"componentmod/internal/utils/log"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/stroiman/go-automapper"
	"gorm.io/gorm"
)

type CarouselService struct {
}

func GetCarouselService() *CarouselService {
	return &CarouselService{}
}

func (r *CarouselService) GetCarouselViewList(p *dto.PageForMultSearchDTO) (interface{}, error) {

	//頁數預設 矯正
	baseApiService := api.GetBaseApiService()
	page, pageLimit := baseApiService.PageParameter(p.Page, p.PageLimit, 1, 15)

	p.Page = page
	p.PageLimit = pageLimit
	carouselData, count, err := r.getCarouselData(p)
	if err != nil {
		return nil, err
	}

	p.Count = count

	data := &backstagedto.CarouselListDTO{
		Carousel: carouselData,
		PageData: p,
	}

	return data, nil
}

func (r *CarouselService) getCarouselData(p *dto.PageForMultSearchDTO) ([]*backstagedto.CarouselData, int64, error) {

	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Carousel{})

	//搜尋條件分類
	if PSearch := p.Search["name"]; PSearch != "" {
		sql = sql.Where("carousels.name LIKE ?", "%"+PSearch+"%")
	}

	//筆數 count
	var count int64 = 0
	sql.Count(&count)

	//分頁 page, pageLimit := pageParameter(p.Page, p.PageLimit, 1, 10)
	sql = sql.Limit(p.PageLimit).Offset((p.Page - 1) * p.PageLimit)

	//排序 依照所選欄位
	baseApiService := api.GetBaseApiService()
	if p.SortColumn == "" || !baseApiService.Check(p.Sort) {
		return nil, 0, utils.CreateApiErr(errorcode.PARAMETER_ERROR_CODE, errorcode.PARAMETER_ERROR)
	}

	if p.SortColumn != "" && baseApiService.Check(p.Sort) {
		sql = sql.Order(fmt.Sprintf("%v %v", p.SortColumn, p.Sort))
	}

	var carouselData []*backstagedto.CarouselData
	sql.Find(&carouselData)

	return carouselData, count, nil
}

func (r *CarouselService) GetCarouselById(id string) (interface{}, error) {

	var carouselData *backstagedto.CarouselData
	var pictureListData []*forestagedto.PictureListData
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Carousel{})
	sql.Find(&carouselData, "carousels.id = ?", id)

	sqlPic := sqldb.Model(&model.Picture{}).Where("pictures.status = ?", true).Order("pictures.weight desc")
	sqlPic = sqlPic.Joins("join carousel_picture on pictures.id=picture_id")
	sqlPic = sqlPic.Joins("join carousels on carousels.id=carousel_id and carousels.id = ?", carouselData.Id)
	sqlPic.Find(&pictureListData)

	for _, v := range pictureListData {
		if v.Name != "" {
			v.PictureUrl = config.WebHost + FILE_PATH + v.Name
		}
	}

	carouselIdDTO := &backstagedto.CarouselIdDTO{
		Carousel: carouselData,
		Picture:  pictureListData,
	}

	return carouselIdDTO, nil
}

func (r *CarouselService) DeleteCarousel(ids []string) (interface{}, error) {

	// 從輪詢圖任務中刪除
	sqldb := db.GetMySqlDB()
	err := sqldb.Transaction(func(tx *gorm.DB) error {
		// 從輪詢任務刪除
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Where("id in ?", ids).Delete(&model.Carousel{}).Error; err != nil {
			// return any error will rollback
			return err
		}

		//從輪詢任務、圖片中繼表單 id搜尋
		var pictureData *[]forestagedto.PictureData
		if err := tx.Table("pictures").Joins("join carousel_picture on picture_id = pictures.id and carousel_id in ?", ids).Distinct("picture_id", "pictures.*").Find(&pictureData).Error; err != nil {
			// return any error will rollback
			return err
		}

		// 從輪詢任務、圖片中繼表單 刪除
		var pictureId []int
		fileRoot := FIXED_FILE_PATH
		for _, v := range *pictureData {
			pictureId = append(pictureId, v.Id)
			if err := file.FileRemove(fileRoot + v.Name); err != nil {
				return err
			}
		}
		// 刪除 中繼表單
		if err := tx.Unscoped().Table("carousel_picture").Where("carousel_id in ?", ids).Delete(&model.Picture{}).Error; err != nil {
			return err
		}
		// 刪除 picture 資料
		if err := tx.Unscoped().Table("pictures").Where("id in ?", pictureId).Delete(&model.Picture{}).Error; err != nil {
			return err
		}

		// // return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		errData := errors.WithMessage(errors.WithStack(err), errorcode.SQL_DELETE_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))
		return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SQL_DELETE_ERROR)
	}

	return nil, nil
}

func (r *CarouselService) CreateCarousel(userInfo *backstagedto.JwtUserInfoDTO, carouselCreateOrEditDTO *backstagedto.CarouselCreateOrEditDTO) (interface{}, error) {

	carousel := model.Carousel{
		Name:         carouselCreateOrEditDTO.Name,
		Weight:       carouselCreateOrEditDTO.Weight,
		Status:       carouselCreateOrEditDTO.Status,
		StartTime:    carouselCreateOrEditDTO.StartTime,
		EndTime:      carouselCreateOrEditDTO.EndTime,
		CreateTime:   time.Now(),
		CreateUserId: userInfo.Id,
	}

	sqldb := db.GetMySqlDB()
	err := sqldb.Transaction(func(tx *gorm.DB) error {

		// 從輪詢圖新增
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Create(&carousel).Error; err != nil {
			// return any error will rollback
			return err
		}

		// 新增picture
		if err := tx.Create(carouselCreateOrEditDTO.Picture).Error; err != nil {
			// return any error will rollback
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		//刪除檔案

		errData := errors.WithMessage(errors.WithStack(err), errorcode.SQL_INSERT_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))

		if strings.Contains(err.Error(), "Duplicate") {
			return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SQL_INSERT_ERROR+": 識別碼(key) 重複,請重新輸入")
		}

		return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SQL_INSERT_ERROR)
	}
	//儲存檔案

	return nil, nil
}

func (r *CarouselService) EditCarousel(userInfo *backstagedto.JwtUserInfoDTO, id string, carouselCreateOrEditDTO *backstagedto.CarouselCreateOrEditDTO) (interface{}, error) {

	var carousel *model.Carousel
	sqldb := db.GetMySqlDB()
	sql := sqldb.Model(&model.Carousel{})
	sql.Where("id = ?", id).Find(&carousel)

	carousel.Name = carouselCreateOrEditDTO.Name
	carousel.Weight = carouselCreateOrEditDTO.Weight
	carousel.Status = carouselCreateOrEditDTO.Status
	carousel.StartTime = carouselCreateOrEditDTO.StartTime
	carousel.EndTime = carouselCreateOrEditDTO.EndTime

	carousel.UpdateTime = time.Now()
	carousel.UpdateUserId = userInfo.Id

	var addPicture []model.Picture
	var updatePicture []model.Picture
	err := sqldb.Transaction(func(tx *gorm.DB) error {

		// 從輪詢圖修改
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		// return any error will rollback

		//修改 carousel
		if err := tx.Save(carousel).Error; err != nil {
			return err
		}

		//刪除 carousel_picture many 2 many
		if err := tx.Table("carousel_picture").Where("carousel_id = ?", id).Delete(carousel).Error; err != nil {
			return err
		}

		var picture []*model.Picture
		if err := tx.Table("pictures").Joins("join carousel_picture on picture_id = pictures.id and carousel_id =?", id).Find(&picture).Error; err != nil {
			return err
		}

		for _, value := range carouselCreateOrEditDTO.Picture {

			//新增 picture array
			if value.Id == 0 {
				pic := model.Picture{}
				automapper.Map(value, &pic)
				//名稱修改
				pic.Name = utils.GetUuidAndTimestamp()
				//存圖片 todo
				utils.SavePicture(FIXED_FILE_PATH+pic.Name, value.PictureUrl)
				addPicture = append(addPicture, pic)
				continue
			}

			for i := 0; i < len(picture); i++ {
				p := picture[i]
				if p.Id != value.Id {
					continue
				}

				//圖片修改
				if value.Name == "" {
					//名稱修改
					p.Name = utils.GetUuidAndTimestamp()
					//圖片儲存 todo
					utils.SavePicture(FIXED_FILE_PATH+p.Name, value.PictureUrl)
				}
				//其他項目修改
				automapper.Map(value, &p)

				// 修改picture
				if err := tx.Save(p).Error; err != nil {
					return err
				} else {
					updatePicture = append(updatePicture, *p)
					picture = append(picture[:i], picture[i+1:]...)
					break
				}
			}
		}

		//批次新增 picture
		if err := tx.Save(addPicture).Error; err != nil {
			return err
		}

		//add picture array
		var addPictureById []int
		for _, v := range addPicture {
			addPictureById = append(addPictureById, v.Id)
		}
		//delete picture array
		var deletePictureById []int
		for _, v := range picture {
			deletePictureById = append(deletePictureById, v.Id)
		}

		//批次刪除 picture
		if err := tx.Delete(&model.Picture{}, deletePictureById).Error; err != nil {
			return err
		}

		//新增 carousel_picture many2many
		var carouselAndPicture []map[string]interface{}
		carouselAndPictureById := append(deletePictureById, addPictureById...)
		for _, v := range carouselAndPictureById {
			cp := map[string]interface{}{"role_id": id, "menu_id": v}
			carouselAndPicture = append(carouselAndPicture, cp)
		}

		if err := tx.Table("carousel_picture").Save(carouselAndPicture).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	//delete picture file
	if err == nil {
		picFile := append(addPicture, updatePicture...)
		for _, v := range picFile {
			filePath := FIXED_FILE_PATH + v.Name
			if file.FileIsExist(filePath) && !file.FileSizeOver(2*file.MBytes, filePath) {
				file.FileRemove(filePath)
			}
		}
	}

	if err != nil {

		errData := errors.WithMessage(errors.WithStack(err), errorcode.SQL_UPDATE_ERROR)
		log.Error(fmt.Sprintf("%+v", errData))

		if strings.Contains(err.Error(), "Duplicate") {
			return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SQL_UPDATE_ERROR)
		}

		return nil, utils.CreateApiErr(errorcode.SERVER_ERROR_CODE, errorcode.SQL_UPDATE_ERROR)
	}

	return nil, nil
}
