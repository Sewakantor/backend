package users

import (
	"capstone/business/users"
	"capstone/controllers"
	"capstone/controllers/users/request"
	"capstone/controllers/users/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.UserUsecaseInterface
}

func NewUserController(UseruseCase users.UserUsecaseInterface) *UserController {
	return &UserController{
		userUseCase: UseruseCase,
	}
}

func (userController *UserController) Register(c echo.Context) error {
	req := request.RegisterUserRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data users.Domain
	data, err = userController.userUseCase.Register(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromUserRegister(data))
}

func (userController *UserController) Login (c echo.Context) error {
	var login users.Domain
	var err error
	var token string
	ctx := c.Request().Context()

	req := request.UserLoginRequest{}
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	login, token, err = userController.userUseCase.Login(ctx, req.Email, req.Password )

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.UserLogin(login, token))
}