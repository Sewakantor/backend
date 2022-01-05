package routes

import (
	"capstone/controllers/addresses"
	"capstone/controllers/complexs"
	"capstone/controllers/facilities"
	"capstone/controllers/roles"
	"capstone/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	JWTMiddleware middleware.JWTConfig
	UserController users.UserController
	RoleController roles.RoleController
	ComplexController complexs.ComplexController
	AddressController addresses.AddressController
	FacilityController facilities.FacilityController
}

func (ctrl *RouteControllerList) RouteRegister(e *echo.Echo) {

	// jwt := middleware.JWTWithConfig(ctrl.JWTMiddleware)

	e.POST("register", ctrl.UserController.Register)
	e.POST("login", ctrl.UserController.Login)

	e.POST("role", ctrl.RoleController.Add)
	e.GET("role", ctrl.RoleController.GetAll)
	e.DELETE("role/:id", ctrl.RoleController.Delete)

	e.POST("complex", ctrl.ComplexController.Add)
	e.GET("complexs", ctrl.ComplexController.GetAll)
	e.GET("complex/:id", ctrl.ComplexController.GetByID)
	e.PUT("complex/:id", ctrl.ComplexController.Update)
	e.DELETE("complex/:id", ctrl.ComplexController.Delete)

	e.POST("address", ctrl.AddressController.Add)
	e.GET("addresses", ctrl.AddressController.GetAll)
	e.PUT("address/:id", ctrl.AddressController.Update)
	e.DELETE("address/:id", ctrl.AddressController.Delete)

	e.POST("facility", ctrl.FacilityController.Add)
	e.GET("facilities", ctrl.FacilityController.GetAll)
	e.PUT("facility/:id", ctrl.FacilityController.Update)
	e.DELETE("facility/:id", ctrl.FacilityController.Delete)

}