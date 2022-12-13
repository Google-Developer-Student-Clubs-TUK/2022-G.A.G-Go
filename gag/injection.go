package main

import (
	"log"

	"gag.com/handler"
	"gag.com/repository"
	"gag.com/service"
	"github.com/gin-gonic/gin"
)

func inject(d *dataSources) (*gin.Engine, error) {
	log.Println("Injecting data sources")

	/*
	 * repository layer
	 */
	userRepository := repository.NewUserRepository(d.DB)
	deviceRepository := repository.NewDeviceRepository(d.DB)
	eclassRepository := repository.NewEclassRepository(d.Eclass)

	/*
	 * service layer
	 */
	userService := service.NewUserService(&service.USConfig{
		UserRepository:   userRepository,
		DeviceRepository: deviceRepository,
		EclassRepository: eclassRepository,
	})

	router := gin.Default()

	handler.NewHandler(&handler.Config{
		R:           router,
		UserService: userService,
	})

	return router, nil
}
