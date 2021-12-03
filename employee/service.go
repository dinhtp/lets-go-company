package employee

import (
    "context"
    "fmt"
    "strconv"

    "gorm.io/gorm"
    pb "github.com/dinhtp/lets-go-pbtype/employee"
    "github.com/gogo/protobuf/types"
)

type Service struct {
    db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
    return &Service{db: db}
}

func (s Service) Get(ctx context.Context, r *pb.OneEmployeeRequest) (*pb.Employee, error) {
    if err := validateOne(r); nil != err {
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())
    fmt.Println(id)
    employee, err := NewRepository(s.db).FindOne(id)
    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(employee), nil
}

func (s Service) Create(ctx context.Context, e *pb.Employee) (*pb.Employee, error) {
    if err := validateCreate(e); nil != err {
        return nil, err
    }

    fmt.Println(e)
    employee, err := NewRepository(s.db).CreatOne(prepareDataToRequest(e))

    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(employee), nil
}

func (s Service) Update(ctx context.Context, r *pb.Employee) (*pb.Employee, error) {
    if err := validateUpdate(r); nil != err {
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())
    employee, err := NewRepository(s.db).UpdateOne(id, prepareDataToRequest(r))

    if nil != err {
        return nil, err
    }

    return prepareDataToResponse(employee), nil
}

func (s Service) List(ctx context.Context, r *pb.ListEmployeeRequest) (*pb.ListEmplyeeResponse, error) {
    var list []*pb.Employee
    if err := validateList(r); nil != err {
        return nil, err
    }
    fmt.Println("In o day")
    fmt.Println(r.GetCompanyid())
    company, count, err := NewRepository(s.db).ListAll(r)
    if nil != err {
        return nil, err
    }
    for i := 0; i < len(company); i++ {
        list = append(list, prepareDataToResponse(company[i]))
    }

    return &pb.ListEmplyeeResponse{
        Items:      list,
        TotalCount: uint32(count),
        Page:       r.GetPage(),
        Limit:      r.GetLimit(),
    }, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneEmployeeRequest) (*types.Empty, error) {
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

