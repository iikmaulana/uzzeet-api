package handler

func (ox gatewayHandler) initRoute() {
	ox.service.Docs("docs/api-docs.yaml")

	ox.service.Strict(
		ox.service.POST("/device", ox.createDevice),
		ox.service.PUT("/device/{imei}", ox.updateDevice),
		ox.service.DELETE("/device/{imei}", ox.deleteDeviceById),
	)
	ox.service.GET("/device/view/{imei}", ox.getDeviceById)
	ox.service.GET("/device/list", ox.getAllDevice)

	ox.service.Strict(
		ox.service.POST("/gpstype", ox.createGpstype),
		ox.service.PUT("/gpstype/{id}", ox.updateGpstype),
		ox.service.DELETE("/gpstype/{id}", ox.deleteGpstypeById),
	)

	ox.service.GET("/gpstype/view/{id}", ox.getGpstypeById)
	ox.service.GET("/gpstype/list", ox.getAllGpstype)

	ox.service.POST("/history", ox.createHisory)
	ox.service.PUT("/history/{id}", ox.updateHisory)
	ox.service.DELETE("/history/{id}", ox.deleteHisoryById)

	ox.service.GET("/history/view/{id}", ox.getHisoryById)
	ox.service.GET("/history/list", ox.getAllHisory)

	ox.service.Strict(
		ox.service.POST("/vehicle", ox.createVehicle),
		ox.service.PUT("/vehicle/{id}", ox.updateVehicle),
		ox.service.DELETE("/vehicle/{id}", ox.deleteVehicleById),
	)

	ox.service.GET("/vehicle/view/{id}", ox.getVehicleById)
	ox.service.GET("/vehicle/list", ox.getAllVehicle)

}
