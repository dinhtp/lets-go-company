package company

import (
    "fmt"
    "github.com/dinhtp/lets-go-company/model"
    pb "github.com/dinhtp/lets-go-pbtype/company"
    "gorm.io/gorm"
    "strings"
    "time"
)

func prepareDataToResponse(c *model.Company) *pb.Company{
    data := &pb.Company{
        Id:                   fmt.Sprintf("%d", c.ID),
        Name:                 c.Name,
        Phone:                c.Phone,
        Email:                c.Email,
        Address:              c.Address,
        TaxNumber:            c.TaxNumber,
        CreatedAt:            c.CreatedAt.Format(time.RFC3339),
        UpdatedAt:            c.UpdatedAt.Format(time.RFC3339),
    }


    // TODO: Calculate total employee

    return data
}

func prepareDataToRequest(p *pb.Company) *model.Company {

    return &model.Company{
        Model:     gorm.Model{},
        Name:      p.Name,
        Phone:     p.Phone,
        Email:     p.Email,
        Address:   p.Address,
        TaxNumber: p.TaxNumber,
    }
}

func getModel(id uint,c *model.Company) *model.Company {
    c.ID = id
    return &model.Company{
        Model:     gorm.Model{
            ID:        id,
            CreatedAt: time.Time{},
            UpdatedAt: time.Time{},
            DeletedAt: gorm.DeletedAt{},
        },
        Name:      c.Name,
        Phone:     c.Phone,
        Email:     c.Email,
        Address:   c.Address,
        TaxNumber: c.TaxNumber,
    }
}

func getList(l []*pb.Company, count uint32,page uint32,limit uint32 ) *pb.ListCompanyResponse{
    return &pb.ListCompanyResponse{
        Items:                l,
        MaxPage:              0,
        TotalCount:           0,
        Page:                 page,
        Limit:                limit,
    }
}

//list []model.company -> list []pb.company
func changeList(l1 [] *model.Company)  ([]*pb.Company){
    var l2 []*pb.Company
    for i := 0; i < len(l1);i++{
        l2 = append(l2,prepareDataToResponse(l1[i]))
    }
    return l2
}

func divideString(str string) []string{
    split := strings.Split(str,",")
    return split
}
