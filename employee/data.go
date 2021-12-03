package employee

import (
    "fmt"
    "gorm.io/gorm"
    "time"
    "strconv"

    "github.com/dinhtp/lets-go-company/model"
    pb "github.com/dinhtp/lets-go-pbtype/employee"
)

func prepareDataToResponse(e *model.Employee) *pb.Employee {
    data := &pb.Employee{
        Id:        fmt.Sprintf("%d", e.ID),
        Companyid: strconv.Itoa(int(e.CompanyId)),
        Name:      e.Name,
        Email:     e.Email,
        Dob:       e.DOB.Format(time.RFC3339),
        Gender:    e.Gender,
        Role:      e.Role,
        CreatedAt: e.CreatedAt.Format(time.RFC3339),
        UpdatedAt: e.UpdatedAt.Format(time.RFC3339),
    }
    return data
}

func prepareDataToRequest(p *pb.Employee) *model.Employee {
    companyid,_ := strconv.Atoi(p.Companyid)
    dob,_ :=  time.Parse(time.RFC3339,p.Dob)
    return &model.Employee{
        Model:     gorm.Model{},
        CompanyId: uint(companyid),
        Name:      p.Name,
        Email:     p.Email,
        DOB:       dob,
        Gender:    p.Gender,
        Role:      p.Role,
    }
}

func getModel(id uint, c *model.Employee) *model.Employee {
    c.ID = id
    return &model.Employee{
        Model: gorm.Model{
            ID:        id,
            CreatedAt: time.Time{},
            UpdatedAt: time.Time{},
            DeletedAt: gorm.DeletedAt{},
        },
        CompanyId: c.CompanyId,
        Name:      c.Name,
        Email:     c.Email,
        DOB:       c.DOB,
        Gender:    c.Gender,
        Role:      c.Role,
    }
}
