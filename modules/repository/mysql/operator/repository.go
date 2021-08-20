package operator

import (
	portOperator "portadapter/business/operator/port"

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

func (db *Repository) CreateData(operator portOperator.SaveOperatorRepo) error {
	if err := db.Create(&operator).Error; err != nil {
		return err
	}
	return nil
}

func (db *Repository) ReadData(ID string) (portOperator.OperatorRepo, error) {
	var operator portOperator.OperatorRepo
	db.Where("id = ?", ID).Find(&operator)
	return operator, nil
}

func (db *Repository) UpdateData(ID string, operator portOperator.SaveOperatorRepo) error {
	if err := db.Model(&portOperator.OperatorRepo{}).Where("id = ?", ID).Update(operator).Error; err != nil {
		return err
	}
	return nil
}

func (db *Repository) DeleteData(ID string) error {
	if err := db.Where("id = ?", ID).Delete(&portOperator.OperatorRepo{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *Repository) ListData(filterLabel string) ([]portOperator.OperatorRepo, error) {
	var operator []portOperator.OperatorRepo
	if filterLabel != "" {
		db.Where("label LIKE ?", "%"+filterLabel+"%").Find(&operator)
	} else {
		db.Find(&operator)
	}
	return operator, nil

}
