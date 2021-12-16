package company

import (
    "fmt"
    "time"

    "github.com/dinhtp/lets-go-company/model"
    pb "github.com/dinhtp/lets-go-pbtype/company"
)

func prepareDataToResponse(c *model.Company) *pb.Company {
    return &pb.Company{
        Id:            fmt.Sprintf("%d", c.ID),
        Name:          c.Name,
        Phone:         c.Phone,
        Email:         c.Email,
        Address:       c.Address,
        TaxNumber:     c.TaxNumber,
        TotalEmployee: 0, // TODO: Calculate total employee
        CreatedAt:     c.CreatedAt.Format(time.RFC3339),
        UpdatedAt:     c.UpdatedAt.Format(time.RFC3339),
    }
}
