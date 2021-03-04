package inventory

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"time"
)

type Endpoints struct {
	InventoryAllocationEndpoint   endpoint.Endpoint
}

func MakeEndpoints(service InventoryService) Endpoints {
	InventoryAllocationEndpoint := makeInventoryAllocationEndpoint(service)

	return Endpoints{
		InventoryAllocationEndpoint:   InventoryAllocationEndpoint,
	}
}

func makeInventoryAllocationEndpoint(service InventoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return StandardJsonResponse{Success: true, Message: "Void Sale", Data: request, Errors: nil, Timestamp: time.Now().String()}, nil
	}
}

func makeVoidSaleEndpoint(service ElavonService) endpoint.Endpoint {
	ep := func(ctx context.Context, UntypedReturnRequest interface{}) (interface{}, error) {
		str, ok := VoidSaleRequest{}, false

		if contextIsCancelled(ctx) {
			return nil, errs.NewContextCancelledError()
		}

		if reflect.TypeOf(&VoidSaleRequest{}) != reflect.TypeOf(UntypedReturnRequest) {

			str, ok = UntypedReturnRequest.(VoidSaleRequest)
			log.Println("Void Sale Request: ", str)

			if !ok {
				log.Println("UntypedVoidSaleRequest: ", UntypedReturnRequest)
				log.Println("reflect.TypeOf(VoidSaleRequest{}): ", reflect.TypeOf(VoidSaleRequest{}))
				log.Println("reflect.TypeOf(UntypedReturnRequest): ", reflect.TypeOf(UntypedReturnRequest))
				return nil, errors.New("unable to cast VoidSaleRequest")
			}
		}

		TransactionData := newTransactionFromVoidSaleRequest(str)
		b, err := service.VoidSale(str.GatewayConfig, TransactionData)

		if err != nil {
			log.Println("makeVoidSaleEndpoint err: ", err)
			return TransactionData, err
		}

		return b, nil
	}

	return ep
}
