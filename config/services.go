package config

import (
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/uzzeet-api/controller"
	"github.com/iikmaulana/uzzeet-api/service/handler"
	"github.com/iikmaulana/uzzeet-api/service/repository/core"
)

func (cfg Config) InitService() serror.SError {

	vehiclesRepo, serr := core.NewVehiclesRepository(cfg.Registry)
	if serr != nil {
		return serr
	}

	deviceRepo, serr := core.NewDeviceRepository(cfg.Registry)
	if serr != nil {
		return serr
	}

	deviceUsecase := controller.NewDeviceUsecase(deviceRepo)
	gpstypeUsecase := controller.NewGpsTypeUsecase(deviceRepo)
	historyUsecase := controller.NewHistoryUsecase(deviceRepo)
	vehicleUsecase := controller.NewVehicleUsecase(vehiclesRepo)

	handler.NewGatewayHandler(cfg.Gateway, deviceUsecase, vehicleUsecase, gpstypeUsecase, historyUsecase)

	return nil
}
