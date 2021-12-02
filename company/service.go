package company

import (
    "context"
    "strconv"

    "github.com/gogo/protobuf/types"
    "gorm.io/gorm"

    pb "github.com/dinhtp/lets-go-pbtype/company"
)

type Service struct {
    db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
   db =  db.Debug()

    return &Service{db: db}
}

func (s Service) Create(ctx context.Context, r *pb.Company) (*pb.Company, error) {
    if err := validateCreate(r); nil != err {
        return nil, err
    }

    company, err := NewRepository(s.db).CreatOne(prepareDataToRequest(r))

    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(company), nil
}

func (s Service) Update(ctx context.Context, r *pb.Company) (*pb.Company, error) {
    if err := validateUpdate(r); nil != err {
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())
    company, err := NewRepository(s.db).UpdateOne(id, prepareDataToRequest(r))

    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(company), nil
}

func (s Service) Get(ctx context.Context, r *pb.OneCompanyRequest) (*pb.Company, error) {
    if err := validateOne(r); nil != err {
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())

    company, err := NewRepository(s.db).FindOne(id)
    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(company), nil
}

func (s Service) List(ctx context.Context, r *pb.ListCompanyRequest) (*pb.ListCompanyResponse, error) {
    var list []*pb.Company
    if err := validateList(r); nil != err {
        return nil, err
    }

    company, count, err := NewRepository(s.db).ListAll(r)

    if nil != err {
        return nil, err
    }
    for i := 0; i < len(company); i++ {
        list = append(list, prepareDataToResponse(company[i]))
    }

    return &pb.ListCompanyResponse{
        Items:      list,
        Page:       r.GetPage(),
        TotalCount: uint32(count),
        Limit:      r.GetLimit(),
    }, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneCompanyRequest) (*types.Empty, error) {
    if err := validateOne(r); nil != err {
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())
    err := NewRepository(s.db).DeleteOne(id)
    if nil != err {
        return nil, err
    }

    return &types.Empty{}, nil
}
