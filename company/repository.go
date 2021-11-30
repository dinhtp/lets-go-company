package company

import (
	"fmt"
	"github.com/dinhtp/lets-go-company/model"
	pb "github.com/dinhtp/lets-go-pbtype/company"
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

func (r *Repository) DeleteOne(id int) (*model.Company, error) {
	var result model.Company

	query := r.db.Model(&model.Company{}).Where("id=?", id).Delete(&result)
	if err := query.Error; nil != err {
		return nil, err
	}

	return &result, nil
}

func (r *Repository) ListAll(rq *pb.ListCompanyRequest) ([]*model.Company,int64, error) {
	var list []*model.Company
	var count int64
	var limit = int(rq.GetLimit())
	var page = int(rq.GetPage())
	var searchFields = rq.GetSearchField()
	var searchValue = rq.GetSearchValue()
	var offset = limit*(page - 1)
	var str = ""

	split := divideString(searchFields)
	if len(searchFields) == 0  {
		query := r.db.Model(&model.Company{}).Limit(limit).Offset(offset).Find(&list)
		if err := query.Error; nil != err {
			return nil,count, err
		}
	}else {
		str = str + split[0] + " like " + "'%" + fmt.Sprintf("%s", searchValue) + "%'"
		//str = fmt.Sprintf("%s like ?",split[0])

		if len(split) > 1 {
			for i := 1; i < len(split); i++ {
				str = str + " OR " + split[i] + " like " + "'%" + fmt.Sprintf("%s", searchValue) + "%'"
			}
		}
		query := r.db.Model(&model.Company{}).Where(str).Limit(limit).Offset(offset).Find(&list)
		if err := query.Error; nil != err {
			return nil, count, err
		}
	}

	//subquery


	//Cau lenh dem
	//subquery := r.db.Model(&model.Company{}).Select("(id)").Count(&count)
	return list,count, nil
}
