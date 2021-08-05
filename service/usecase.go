package service

import (
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/uzzeet-api/models"
)

type VehicleUsecase interface {
	AddVehicleUsecase(organizationID string, form models.VehVehiclesRequest) (serr serror.SError)
	UpdateVehicleUsecase(id string, form models.VehVehiclesUpdate) (serr serror.SError)
	DeleteVehicleByIdUsecase(id string) (serr serror.SError)
	GetVehicleByIdUsecase(id string) (result []models.VehVehiclesResult, serr serror.SError)
	GetAllVehicleUsecase() (result models.VehListVehiclesResult, serr serror.SError)
}

type DeviceUsecase interface {
	AddDeviceUsecase(form models.DeviceRequest) (serr serror.SError)
	UpdateDeviceUsecase(imei string, form models.UpdateDeviceRequest) (serr serror.SError)
	GetDeviceByImeiUsecase(imei string) (result []models.DeviceResult, serr serror.SError)
	GetAllDeviceUsecase() (result models.ListDeviceResult, serr serror.SError)
	DeleteDeviceByImeiUsecase(imei string) (serr serror.SError)
}

type GpsTypeUsecase interface {
	AddGpsTypeUsecase(form models.GpsTypeRequest) (serr serror.SError)
	UpdateGpsTypeUsecase(id int64, form models.UpdateGpsTypeRequest) (serr serror.SError)
	GetGpsTypeByIDUsecase(id int64) (result []models.GpsTypeResult, serr serror.SError)
	GetAllGpsTypeUsecase(ndata int64, page int) (result models.ListGpsTypeResult, serr serror.SError)
	DeleteGpsTypeIdUsecase(id int64) (serr serror.SError)
}

type HistoryUsecase interface {
	AddHistoryUsecase(form models.HistoryRequest) (serr serror.SError)
	UpdateHistoryUsecase(id string, form models.UpdateHistoryRequest) (serr serror.SError)
	GetHistoryByIDUsecase(id string) (result []models.HistoryResult, serr serror.SError)
	GetAllHistoryUsecase(ndata int64, page int) (result models.ListHistoryResult, serr serror.SError)
	DeleteHistoryByIdUsecase(id string) (serr serror.SError)
}
