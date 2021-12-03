package company

import (
    "fmt"
    "strings"

    "gorm.io/gorm"

    "github.com/dinhtp/lets-go-company/model"
    pb "github.com/dinhtp/lets-go-pbtype/company"
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

func (r *Repository) CreatOne(c *model.Company) (*model.Company, error) {
    query := r.db.Model(&model.Company{}).Create(c)

    if err := query.Error; nil != err {
        return nil, err
    }

    return c, nil
}

func (r *Repository) UpdateOne(id int, c *model.Company) (*model.Company, error) {
    query := r.db.Model(&model.Company{}).Where("id=?", id).UpdateColumns(getModel(uint(id), c))

    if err := query.Error; nil != err {
        return nil, err
    }

    return c, nil
}

func (r *Repository) DeleteOne(id int) error {
    var result model.Company

    query := r.db.Model(&model.Company{}).Where("id=?", id).Delete(&result)
    if err := query.Error; nil != err {
        return err
    }

    return nil
}

func (r *Repository) ListAll(req *pb.ListCompanyRequest) ([]*model.Company, int64, error) {
    var count int64
    var list []*model.Company

    sql := ""
    limit := int(req.GetLimit())
    offset := limit * int(req.GetPage()-1)

    if req.GetSearchField() != "" && req.GetSearchValue() != "" {
        searchFields := strings.Split(req.GetSearchField(), ",")
        searchValue := fmt.Sprintf("'%%%s%%'", req.GetSearchValue())

        for idx, field := range searchFields {
            if idx == 0 {
                sql += fmt.Sprintf("%s LIKE %s", field, searchValue)
                continue
            }
            sql += fmt.Sprintf(" OR %s LIKE %s", field, searchValue)
        }
    }

    listQuery := r.db.Model(&model.Company{}).Select("*")
    countQuery := r.db.Model(&model.Company{}).Select("id")

    if sql != "" {
        countQuery = countQuery.Where(sql)
        listQuery = listQuery.Where(sql)
    }

    if err := countQuery.Count(&count).Error; nil != err {
        return nil, 0, err
    }

    if err := listQuery.Find(&list).Limit(limit).Offset(offset).Error; nil != err {
        return nil, 0, err
    }

    return list, count, nil
}
