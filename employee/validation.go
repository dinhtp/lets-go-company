package employee

import (
    pb "github.com/dinhtp/lets-go-pbtype/employee"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "strings"
)

func validateOne(r *pb.OneEmployeeRequest) error {
    if "" == r.GetId() {
        return status.Error(codes.InvalidArgument, "company id is required")
    }
    return nil
}

func validateUpdate(e *pb.Employee) error {
    if "" == e.GetId() {
        return status.Error(codes.InvalidArgument, "CompanyID is required")
    }
    return validateCreate(e)
}

func validateCreate(e *pb.Employee) error {

    if e.GetName() == "" {
        return status.Error(codes.InvalidArgument, "Name is required")
    }
    if e.GetEmail() == "" {
        return status.Error(codes.InvalidArgument, "Email is required")
    }
    if e.GetDob() == "" {
        return status.Error(codes.InvalidArgument, "Dob is required")
    }
    if e.GetGender() == "" {
        return status.Error(codes.InvalidArgument, "Gender is required")
    }
    if e.GetRole() == "" {
        return status.Error(codes.InvalidArgument, "Role is required")
    }
    return nil
}

func validateList(e *pb.ListEmployeeRequest) error  {
    var field = strings.Split(e.GetSearchField(), ",")

    if e.GetPage() <= 0 {
        return status.Error(codes.InvalidArgument, "Invalid Page")
    }

    if e.GetLimit() <= 0 {
        return status.Error(codes.InvalidArgument, "Invalid Limit")
    }

    if  e.GetSearchField() == "" && e.GetSearchValue() ==""{
        return nil
    }

    for i := 0; i < len(field); i++ {
        var newfield = strings.ToLower(strings.TrimSpace(field[i]))
        if newfield != "name" && newfield != "dob" && newfield != "email" && newfield != "gender" && newfield != "role" && newfield != "" {
            return status.Error(codes.InvalidArgument, "Invalid SearchFields")
        }
    }
    return nil
}