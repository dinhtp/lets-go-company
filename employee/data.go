package employee


import (
	"fmt"
	"time"

	"github.com/dinhtp/lets-go-company/model"
	pb "github.com/dinhtp/lets-go-pbtype/company"
)

func prepareDataToResponse(e *model.Employee) *pb.Employee{
	data := &pb.Employee{
		Id:                   fmt.Sprintf("%d", e.ID),
		Name:                 e.Name,
		Phone:                e.Phone,
		Email:                e.Email,
		DOB:                  e.DOB,
		Gender:               e.Gender,
		Role:                 e.Role,
		CreatedAt:            e.CreatedAt.Format(time.RFC3339),
		UpdatedAt:            e.UpdatedAt.Format(time.RFC3339),
	}

	// TODO: Calculate total employee

	return data
}