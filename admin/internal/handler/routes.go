// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/iot-platform/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/device/list",
				Handler: DeviceListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/device/create",
				Handler: DeviceCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/device/modify",
				Handler: DeviceModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/device/delete",
				Handler: DeviceDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/product/list",
				Handler: ProductListHandler(serverCtx),
			},
		},
	)
}
