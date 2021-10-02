package company

import (
    "github.com/dinhtp/lets-go-company/model"
    "gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) FindOne(id int) (*model.Company, error) {
    var result model.Company

    query := r.db.Model(&model.Company{}).Where("id = ?", id).First(&result)

    if err := query.Error; nil != err {
        return nil, err
    }

    return &result, nil
}
