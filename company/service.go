package company

import (
    "context"
    "gorm.io/gorm"

    "github.com/gogo/protobuf/types"

    pb "github.com/dinhtp/lets-go-pbtype/company"

)

type Service struct {
    db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
    return &Service{db: db}
}

func (s Service) Create(ctx context.Context, r *pb.Company) (*pb.Company, error) {
    return &pb.Company{}, nil
}

func (s Service) Update(ctx context.Context, r *pb.Company) (*pb.Company, error) {
    return &pb.Company{}, nil
}

func (s Service) Get(ctx context.Context, r *pb.OneCompanyRequest) (*pb.Company, error) {
    return &pb.Company{}, nil
}

func (s Service) List(ctx context.Context, r *pb.ListCompanyRequest) (*pb.ListCompanyResponse, error) {
    return &pb.ListCompanyResponse{}, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneCompanyRequest) (*types.Empty, error) {
    return &types.Empty{}, nil
}
