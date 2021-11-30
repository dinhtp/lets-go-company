package company

import (
    "context"
    "gorm.io/gorm"
    "strconv"

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
    if err := validateCreate(r); nil!=err {
        return nil,err
    }

    company,err := NewRepository(s.db).CreatOne(prepareDataToRequest(r))

    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(company), nil
}

func (s Service) Update(ctx context.Context, r *pb.Company) (*pb.Company, error) {
    if err := validateUpdate(r); nil != err {
        return nil, err
    }

    id , _ := strconv.Atoi(r.GetId())
    company,err := NewRepository(s.db).UpdateOne(id , prepareDataToRequest(r))

    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(company), nil
}

func (s Service) Get(ctx context.Context, r *pb.OneCompanyRequest) (*pb.Company, error) {
    if err := validateOne(r); nil != err {
        return nil, err
    }

    id , _ := strconv.Atoi(r.GetId())

    company, err := NewRepository(s.db).FindOne(id)
    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(company), nil
}

func (s Service) List(ctx context.Context, r *pb.ListCompanyRequest) (*pb.ListCompanyResponse, error) {
    if err := validateList(r); nil != err {
        return nil, err
    }

    company, count, err := NewRepository(s.db).ListAll(r)
    //chuyen tu []model.company -> list []pb.company
    list := changeList(company)
    if nil != err {
        return nil, err
    }

    return getList(list,uint32(count),r.GetPage(),r.GetLimit()), nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneCompanyRequest) (*types.Empty, error) {
    if err := validateOne(r); nil != err {
        return nil, err
    }

    id , _ := strconv.Atoi(r.GetId())
    NewRepository(s.db).DeleteOne(id)

    return &types.Empty{}, nil
}
