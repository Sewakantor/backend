package routes

import (
	"capstone/controllers/buildings"
	"capstone/controllers/complexs"
	"capstone/controllers/facilities"
	"capstone/controllers/roles"
	"capstone/controllers/units"
	"capstone/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	JWTMiddleware 		middleware.JWTConfig
	UserController 		users.UserController
	RoleController 		roles.RoleController
	ComplexController 	complexs.ComplexController
	FacilityController 	facilities.FacilityController
	BuildingController 	buildings.BuildingController
	UnitController 		units.UnitController
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

	e.POST("facility", ctrl.FacilityController.Add)
	e.GET("facilities", ctrl.FacilityController.GetAll)
	e.PUT("facility/:id", ctrl.FacilityController.Update)
	e.DELETE("facility/:id", ctrl.FacilityController.Delete)

	e.POST("building", ctrl.BuildingController.Add)
	e.GET("buildings", ctrl.BuildingController.GetAll)
	e.GET("building/:id", ctrl.BuildingController.GetByID)
	e.GET("building/complex/:complexid", ctrl.BuildingController.GetByComplexID)
	e.PUT("building/:id", ctrl.BuildingController.Update)
	e.DELETE("building/:id", ctrl.BuildingController.Delete)

	e.POST("unit", ctrl.UnitController.Add)
	e.GET("units", ctrl.UnitController.GetAll)
	e.GET("unit/:id", ctrl.UnitController.GetByID)
	e.PUT("unit/:id", ctrl.UnitController.Update)
	e.DELETE("unit/:id", ctrl.UnitController.Delete)

}