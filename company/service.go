package company

import (
    "context"
    "math"
    "strconv"

    "github.com/gogo/protobuf/types"
    "gorm.io/gorm"

    pb "github.com/dinhtp/lets-go-pbtype/company"
    "github.com/dinhtp/lets-go-company/model"
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

    company, err := NewRepository(s.db).FindOne(id)
    mapEmployee, err := NewRepository(s.db).countTotalEmployee(id)

    companyData := prepareDataToResponse(company)
    companyData.TotalEmployee = mapEmployee[uint(id)]

    if nil != err {
        return nil, err
    }

    return companyData, nil
}

func (s Service) List(ctx context.Context, r *pb.ListCompanyRequest) (*pb.ListCompanyResponse, error) {
    if err := validateList(r); nil != err {
        return nil, err
    }

    var list []*pb.Company
    companyChanel := make(chan []*model.Company, 1)
    countChanel := make(chan int64, 1)
    mapChanel := make(chan map[uint]uint32, 1)
    errorChanel := make(chan error, 2)

    go func() {
        company, count, err := NewRepository(s.db).ListAll(r)
        errorChanel <- err
        companyChanel <- company
        countChanel <- count
    }()

    go func() {
        mapEmployee, err := NewRepository(s.db).countTotalEmployee(0)
        errorChanel <- err
        mapChanel <- mapEmployee
    }()

    company := <-companyChanel
    count := <-countChanel
    mapEmployee := <-mapChanel
    err := <-errorChanel
    if nil != err {
        return nil, err
    }
    err2 := <-errorChanel
    if nil != err2 {
        return nil, err2
    }

    for i := range company {
        companyData := prepareDataToResponse(company[i])
        companyData.TotalEmployee = mapEmployee[company[i].ID]
        list = append(list, companyData)
    }

    return &pb.ListCompanyResponse{
        Items:      list,
        Page:       r.GetPage(),
        TotalCount: uint32(count),
        Limit:      r.GetLimit(),
        MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
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
