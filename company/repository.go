package company

import (
    "fmt"
    "github.com/dinhtp/lets-go-company/model"
    "gorm.io/gorm"
    "strings"

    pb "github.com/dinhtp/lets-go-pbtype/company"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    db = db.Debug()

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

    // assign value, pre process
    sql := ""
    limit := int(req.GetLimit())
    offset := limit * int(req.GetPage()-1)
    searchFields := strings.Split(req.GetSearchField(), ",")
    searchValue := fmt.Sprintf("'%%%s%%'", req.GetSearchValue())

    // build sql query in string format
    for idx, field := range searchFields {
        if idx == 0  {
            sql += fmt.Sprintf("%s LIKE %s", field, searchValue)
            continue
        }
        sql += fmt.Sprintf(" OR %s LIKE %s", field, searchValue)
    }

    // build base query
    listQuery := r.db.Model(&model.Company{}).Select("*")
    countQuery := r.db.Model(&model.Company{}).Select("id")

    // apply filter
    if req.GetSearchField() != "" && req.GetSearchValue() != "" {
        countQuery = countQuery.Where(sql)
        listQuery = listQuery.Where(sql)
    }

    // count records and check error
    if err := countQuery.Count(&count).Error; nil != err {
        return nil, 0, err
    }

    // list records and check error
    if err := listQuery.Find(&list).Limit(limit).Offset(offset).Error; nil != err {
        return nil, 0, err
    }

    //var
    //var page = int(rq.GetPage())
    //var searchFields = rq.GetSearchField()
    //var searchValue = rq.GetSearchValue()
    //var
    //var str = ""
    //var query *gorm.DB
    //var subquery *gorm.DB
    //
    //
    //
    //if searchFields == "" {
    //    query := r.db.Model(&model.Company{}).Limit(limit).Offset(offset).Find(&list)
    //    subquery := r.db.Model(&model.Company{}).Select("id").Count(&count)
    //    if err := query.Error; nil != err {
    //        return nil, 0, err
    //    }
    //    if err := subquery.Error; nil != err {
    //        return nil, 0, err
    //    }
    //    return list, count, nil
    //}
    //

    //query.Where(str)
    //subquery.Where(str)
    ////query := r.db.Model(&model.Company{}).Where(str).Limit(limit).Offset(offset).Find(&list)
    //if err := query.Error; nil != err {
    //    return nil, 0, err
    //}
    //if err := subquery.Error; nil != err {
    //    return nil, 0, err
    //}
    return list, count, nil
}
