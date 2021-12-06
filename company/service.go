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

    company, totalEmployee, err := NewRepository(s.db).FindOne(id)
    if nil != err {
        return nil, err
    }

    return prepareDataToResponse2(company, totalEmployee), nil
}

func (s Service) List(ctx context.Context, r *pb.ListCompanyRequest) (*pb.ListCompanyResponse, error) {
    var list []*pb.Company
    var maxPage uint32
    if err := validateList(r); nil != err {
        return nil, err
    }

    company, count, err := NewRepository(s.db).ListAll(r)
    mapEmployee, err := NewRepository(s.db).countTotalEmployee()
    if nil != err {
        return nil, err
    }

    for i := 0; i < len(company); i++ {
        companyData := prepareDataToResponse2(company[i], mapEmployee[company[i].ID])
        list = append(list, companyData)
    }

    maxPage = uint32(count) / r.GetLimit()
    if uint32(count)%r.GetLimit() > 0 {
        maxPage = (uint32(count) / r.GetLimit()) + 1
    }

    return &pb.ListCompanyResponse{
        Items:      list,
        Page:       r.GetPage(),
        TotalCount: uint32(count),
        Limit:      r.GetLimit(),
        MaxPage:    maxPage,
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
