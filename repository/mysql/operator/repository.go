package operator

import (
	structRespository "portadapter/struct/repository"

	"github.com/jinzhu/gorm"
)

type (
	Repository struct {
		*gorm.DB
	}
)

func New(db *gorm.DB) *Repository {
	return &Repository{
		db.Table("operator"),
	}
}

func (db *Repository) CreateData(operator structRespository.SaveOperator) error {
	if err := db.Create(&operator).Error; err != nil {
		return err
	}
	return nil
}

func (db *Repository) ReadData(ID string) (structRespository.Operator, error) {
	var operator structRespository.Operator
	db.Where("id = ?", ID).Find(&operator)
	return operator, nil
}

func (db *Repository) UpdateData(ID string, operator structRespository.SaveOperator) error {
	if err := db.Model(&structRespository.Operator{}).Where("id = ?", ID).Update(operator).Error; err != nil {
		return err
	}
	return nil
}

func (db *Repository) DeleteData(ID string) error {
	if err := db.Where("id = ?", ID).Delete(&structRespository.Operator{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *Repository) ListData(filterLabel string) ([]structRespository.Operator, error) {
	var operator []structRespository.Operator
	if filterLabel != "" {
		db.Where("label LIKE ?", "%"+filterLabel+"%").Find(&operator)
	} else {
		db.Find(&operator)
	}
	return operator, nil

}
