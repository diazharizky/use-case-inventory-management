package errors

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewInternalError(errMeta map[string]string) error {
	st := status.New(codes.Internal, "Internal error")

	errDetails := &errdetails.ErrorInfo{}
	errDetails.Metadata = errMeta

	st, err := st.WithDetails(errDetails)
	if err != nil {
		panic(
			fmt.Sprintf("Unexpected error attaching metadata: %v", err),
		)
	}

	return st.Err()
}
