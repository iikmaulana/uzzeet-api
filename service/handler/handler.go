package handler

import (
	gateway "github.com/iikmaulana/gateway/service"
	"github.com/iikmaulana/uzzeet-api/service"
)

type gatewayHandler struct {
	service        *gateway.Service
	deviceUsecase  service.DeviceUsecase
	vehicleUsecase service.VehicleUsecase
	gpstypeUsecase service.GpsTypeUsecase
	historyUsecase service.HistoryUsecase
}

func NewGatewayHandler(svc *gateway.Service,
	deviceUsecase service.DeviceUsecase,
	vehicleUsecase service.VehicleUsecase,
	gpstypeUsecase service.GpsTypeUsecase,
	historyUsecase service.HistoryUsecase,
) {
	h := gatewayHandler{
		service:        svc,
		deviceUsecase:  deviceUsecase,
		vehicleUsecase: vehicleUsecase,
		gpstypeUsecase: gpstypeUsecase,
		historyUsecase: historyUsecase,
	}

	h.initRoute()
}
