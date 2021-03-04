package inventory

import (
	"github.com/go-kit/kit/endpoint"
	kitHTTP "github.com/go-kit/kit/transport/http"
	kitHttpTransport "github.com/go-kit/kit/transport/http"
)

func InventoryAllocatorHandler(ep endpoint.Endpoint, serverOptions []kitHTTP.ServerOption) *kitHttpTransport.Server {
	return kitHttpTransport.NewServer(
		ep,
		decodeHTTPInventoryAllocatorRequest,
		encodeInventoryAllocatorHTTPResponse,
		serverOptions...,
	)
}