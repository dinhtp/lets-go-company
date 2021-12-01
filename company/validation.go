package company

import (
    "strings"

    pb "github.com/dinhtp/lets-go-pbtype/company"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func validateOne(r *pb.OneCompanyRequest) error {
    if "" == r.GetId() {
        return status.Error(codes.InvalidArgument, "company id is required")
    }
    return nil
}

func validateCreate(r *pb.Company) error {
    if r.GetName() == "" || r.GetPhone() == "" || r.GetEmail() == "" || r.GetAddress() == "" || r.GetTaxNumber() == "" {
        return status.Error(codes.InvalidArgument, "Attributed is required")
    }
    return nil
}

func validateUpdate(r *pb.Company) error {
    if "" == r.GetId() {
        return status.Error(codes.InvalidArgument, "company id is required")
    }
    return validateCreate(r)
}

func validateList(r *pb.ListCompanyRequest) error {
    var page = r.GetPage()
    var limit = r.GetLimit()
    var field = strings.Split(r.GetSearchField(), ",")

    if (page <= 0) || (limit <= 0) {
        return status.Error(codes.InvalidArgument, "Invalid Attribute")
    }

    for i := 0; i < len(field); i++ {
        var newfield = strings.ToLower(strings.TrimSpace(field[i]))
        if newfield != "name" && newfield != "phone" && newfield != "email" && newfield != "address" && newfield != "tax_number" && newfield != "" {
            return status.Error(codes.InvalidArgument, "Invalid SearchFields")
        }
    }
    return nil
}
