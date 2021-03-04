package inventory

func encodeInventoryAllocatorHTTPResponse(ctx context.Context, rw http.ResponseWriter, response interface{}) error {
	transType := "Inventory Allocator"

	encodeHTTPResponse(ctx, rw, response, transType)

	return nil
}

func encodeHTTPResponse(ctx context.Context, rw http.ResponseWriter, response interface{}, transType string) error {
	log.Println("encodeHTTPResponse start")

	if contextIsCancelled(ctx) {
		return errs.NewContextCancelledError()
	}

	if svcErr, isSvcErr := response.(error); isSvcErr {
		errs.EncodeJSONError(nil, svcErr, rw)
		return nil
	}

	w.Header().Set(constants.ContentHeader, constants.JsonContent)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpStatus)

	//bytes, _ := json.Marshal(standardJsonResponse)

	w.Write(bytes)
}