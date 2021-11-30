package company

import (
	pb "github.com/dinhtp/lets-go-pbtype/company"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func validateOne(r *pb.OneCompanyRequest) error {
	if "" == r.GetId() {
		return status.Error(codes.InvalidArgument, "company id is required")
	}

	return nil
}

func validateCreate(r *pb.Company) error {
	//TODO: validate create
	if (r.GetName() == "") || (r.GetPhone() == "") || (r.GetEmail() == "") || (r.GetAddress() == "") || (r.GetTaxNumber() == "") {
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
	var str = strings.Split(r.GetSearchField(), ",")

	if (page < 0) || (limit < 0) {
		return status.Error(codes.InvalidArgument, "Invalid Attribute")
	}

	for i := 0; i < len(str); i++ {
		var strnw = strings.ToLower(strings.TrimSpace(str[i]))
		if strnw != "name" && strnw != "phone" && strnw != "email" && strnw != "address" && strnw != "tax_number" && strnw != "" {
			return status.Error(codes.InvalidArgument, "Invalid SearchFields")
		}
	}
	return nil
}
