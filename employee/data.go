package employee

import (
    "fmt"
    "time"

    "github.com/dinhtp/lets-go-company/model"
    pb "github.com/dinhtp/lets-go-pbtype/employee"
)

func prepareDataToResponse(e *model.Employee) *pb.Employee {
    data := &pb.Employee{
        Id:        fmt.Sprintf("%d", e.ID),
        Name:      e.Name,
        Email:     e.Email,
        Dob:       fmt.Sprintf("%s", e.DOB),
        Gender:    e.Gender,
        Role:      e.Role,
        CreatedAt: e.CreatedAt.Format(time.RFC3339),
        UpdatedAt: e.UpdatedAt.Format(time.RFC3339),
    }
    return data
}
