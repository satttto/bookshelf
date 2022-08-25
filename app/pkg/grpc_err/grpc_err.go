package grpcerr

// type Code uint32

// // Codes: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
// // TODO: delete this after writing handlers for all the codes.
// const (
// 	OK                 Code = 0  // 200
// 	Canceled           Code = 1  // 499
// 	Unknown            Code = 2  // 500
// 	InvalidArgument    Code = 3  // 400 (Bad Request)
// 	DeadlineExceeded   Code = 4  // 504 (Gateway Timeout)
// 	NotFound           Code = 5  // 404
// 	AlreadyExists      Code = 6  // 409 (Conflict)
// 	PermissionDenied   Code = 7  // 403
// 	ResourceExhausted  Code = 8  // 429
// 	FailedPrecondition Code = 9  // 400
// 	Aborted            Code = 10 // 409
// 	OutOfRange         Code = 11 // 400
// 	Unimplemented      Code = 12 // 501
// 	Internal           Code = 13 // 500
// 	Unavailable        Code = 14 // 503
// 	DataLoss           Code = 15 // 500
// 	Unauthenticated    Code = 16 // 401
// )

// convertErr takes an error value as input and return a corresponding gRPC error
func convertErr(err error) error {
	if err == nil {
		return nil
	}

	return nil
	//status, ok := status.FromError(err)
	//if !ok {
	//	return status.New(codes.Unknown, "").Err()
	//}
}
