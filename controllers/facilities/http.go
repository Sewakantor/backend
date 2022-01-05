package facilities

import (
	"capstone/business/facilities"
	"capstone/controllers"
	"capstone/controllers/facilities/request"
	"capstone/controllers/facilities/response"
	"capstone/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FacilityController struct {
	facilityUseCase facilities.FacilityUsecaseInterface
}

func NewFacilityController(FacilityUseCase facilities.FacilityUsecaseInterface) *FacilityController {
	return &FacilityController{
		facilityUseCase: FacilityUseCase,
	}
}

func (facilityController *FacilityController) Add(c echo.Context) error {
	req := request.FacilityRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	var data facilities.Domain
	data, err = facilityController.facilityUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainFacility(data))
}

func (facilityController *FacilityController) GetAll(c echo.Context) error {
	req := c.Request().Context()
	facility, err := facilityController.facilityUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainFacilityArray(facility))
}

func (facilityController *FacilityController) Update(c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.FacilityRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := facilityController.facilityUseCase.Update(ctx, convID, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainFacility(data))

}

func (facilityController *FacilityController) Delete(c echo.Context) error {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = facilityController.facilityUseCase.Delete(ctx, convID)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}