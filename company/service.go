package company

import (
    "context"
    "math"
    "strconv"

    "github.com/gogo/protobuf/types"
    "gorm.io/gorm"

    "github.com/dinhtp/lets-go-company/model"
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

    var company *model.Company
    var mapEmployee map[uint]uint32

    companyChanel := make(chan *model.Company, 1)
    mapChanel := make(chan map[uint]uint32, 1)
    errorChanel := make(chan error, 2)

    id, _ := strconv.Atoi(r.GetId())

    go func() {
        company, err := NewRepository(s.db).FindOne(id)
        companyChanel <- company
        errorChanel <- err
    }()

    go func() {
        mapEmployee, err := NewRepository(s.db).countTotalEmployee(id)
        mapChanel <- mapEmployee
        errorChanel <- err
    }()

    for i := range errorChanel {
        if i != nil {
            company = nil
            mapEmployee = nil
        } else {
            company = <-companyChanel
            mapEmployee = <-mapChanel
        }
    }

    companyData := prepareDataToResponse(company)
    companyData.TotalEmployee = mapEmployee[uint(id)]

    return companyData, nil
}

func (s Service) List(ctx context.Context, r *pb.ListCompanyRequest) (*pb.ListCompanyResponse, error) {
    if err := validateList(r); nil != err {
        return nil, err
    }

    var list []*pb.Company
    var company []*model.Company
    var count int64
    var mapEmployee map[uint]uint32

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

    for i := range errorChanel {
        if i != nil {
            company = nil
            count = 0
            mapEmployee = nil
        } else {
            company = <-companyChanel
            count = <-countChanel
            mapEmployee = <-mapChanel
        }
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
