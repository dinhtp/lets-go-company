package company

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/dinhtp/lets-go-pbtype/company"

)

func validateOne(r *pb.OneCompanyRequest) error {
	if "" == r.GetId() {
		return status.Error(codes.InvalidArgument, "company id is required")
	}

	return nil
}
