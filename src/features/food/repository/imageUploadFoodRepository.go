package repository

import (
	"context"
	_interface "main/features/food/model/interface"
	"main/utils"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func NewImageUploadFoodRepository(gormDB *gorm.DB) _interface.IImageUploadFoodRepository {
	return &ImageUploadFoodRepository{GormDB: gormDB}
}

func (g *ImageUploadFoodRepository) FindOneAndUpdateFoodImages(ctx context.Context, foodID uint, filename string) error {
	if err := mysql.GormMysqlDB.WithContext(ctx).Model(&mysql.FoodImages{}).Where("id = ?", foodID).Update("image", filename).Error; err != nil {
		return utils.ErrorMsg(ctx, utils.ErrInternalDB, utils.Trace(), utils.HandleError(err.Error(), foodID, filename), utils.ErrFromMysqlDB)
	}
	return nil
}
