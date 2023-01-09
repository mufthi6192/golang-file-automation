package FileRepository

import (
	"fileAutomation/app/entity/model"
	"gorm.io/gorm"
)

type fileRepository struct {
	DB *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileInterface {
	return &fileRepository{DB: db}
}

func (repository *fileRepository) InsertFileInterface(models *model.ProductModels) (string, error) {
	err := repository.DB.Create(models).Error
	if err != nil {
		return "Failed insert database", err
	}
	return "Successfully insert database", nil
}
