package routes

import (
	"portadapter/config"

	operatorController "portadapter/api/intl/v1/operator"
	operatorRepository "portadapter/repository/mysql/operator"
	operatorService "portadapter/service/operator"

	"github.com/labstack/echo"
)

func API(e *echo.Echo) {

	// Init Repository
	db := config.DB
	operatorRepository := operatorRepository.New(db)

	// Init Service
	operatorService := operatorService.New(operatorRepository)

	// Init Controller
	operatorController := operatorController.New(operatorService)
	operator := e.Group("/api/v1/operator")
	operator.POST("", operatorController.CreateData)
	operator.GET("/:id", operatorController.ReadData)
	operator.PUT("/:id", operatorController.UpdateData)
	operator.DELETE("/:id", operatorController.DeleteData)
	operator.GET("", operatorController.ListData)
}
