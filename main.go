package main

import (
	"github.com/Zavr22/car-speed-control/config"
	"github.com/Zavr22/car-speed-control/internal/controller"
	"github.com/Zavr22/car-speed-control/internal/repository"
	"github.com/Zavr22/car-speed-control/internal/service"
	"github.com/Zavr22/car-speed-control/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	repo := repository.NewSpeedRepository("db.csv")
	srv := service.NewSpeedService(repo)
	controller := controller.NewSpeedController(srv, cfg)

	r.POST("/speed", controller.AddSpeedRecord)
	r.GET("/speed/records", middleware.AccessTimeMiddleware(), controller.GetSpeedRecords)
	r.GET("/speed/stats", middleware.AccessTimeMiddleware(), controller.GetSpeedStats)

	r.Run(":8080")
}
