package addresses

import (
	"capstone/business/addresses"
	"capstone/controllers"
	"capstone/controllers/addresses/request"
	"capstone/controllers/addresses/response"
	"capstone/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddressController struct {
	addressUseCase addresses.AddressUsecaseInterface
}

func NewAddressController(AddressUseCase addresses.AddressUsecaseInterface) *AddressController {
	return &AddressController{
		addressUseCase: AddressUseCase,
	}
}

func (addressController *AddressController) Add(c echo.Context) error {
	req := request.AddressRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	var data addresses.Domain
	data, err = addressController.addressUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainAddress(data))
}

func (addressController *AddressController) GetAll(c echo.Context) error {
	req := c.Request().Context()
	address, err := addressController.addressUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainAddressArray(address))
}

func (addressController *AddressController) Update(c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.AddressRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := addressController.addressUseCase.Update(ctx, convID, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainAddress(data))

}


func (addressController *AddressController) Delete(c echo.Context) error {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = addressController.addressUseCase.Delete(ctx, convID)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}