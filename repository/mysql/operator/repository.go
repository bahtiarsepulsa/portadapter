package operator

import (
	structRepository "portadapter/struct/repository"

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

func (db *Repository) CreateData(operator structRepository.SaveOperator) error {
	if err := db.Create(&operator).Error; err != nil {
		return err
	}
	return nil
}

func (db *Repository) ReadData(ID string) (structRepository.Operator, error) {
	var operator structRepository.Operator
	db.Where("id = ?", ID).Find(&operator)
	return operator, nil
}

func (db *Repository) UpdateData(ID string, operator structRepository.SaveOperator) error {
	if err := db.Model(&structRepository.Operator{}).Where("id = ?", ID).Update(operator).Error; err != nil {
		return err
	}
	return nil
}

func (db *Repository) DeleteData(ID string) error {
	if err := db.Where("id = ?", ID).Delete(&structRepository.Operator{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *Repository) ListData(filterLabel string) ([]structRepository.Operator, error) {
	var operator []structRepository.Operator
	if filterLabel != "" {
		db.Where("label LIKE ?", "%"+filterLabel+"%").Find(&operator)
	} else {
		db.Find(&operator)
	}
	return operator, nil

}
