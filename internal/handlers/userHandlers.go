package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/shockwxve/restapi/internal/userService"
	"github.com/shockwxve/restapi/internal/web/users"
	"gorm.io/gorm"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (uh *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := uh.Service.ReadAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		u := users.User{
			Id:    &usr.ID,
			Email: &usr.Email,
		}
		response = append(response, u)
	}

	return response, nil
}

func (uh *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	if userRequest.Email == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email is required")
	}
	if userRequest.Password == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Password is required")
	}

	userToCreate := userService.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	createdUser, err := uh.Service.CreateUser(userToCreate)
	if err != nil {
		// Проверка на дубликат email
		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)` {
			return nil, echo.NewHTTPError(http.StatusConflict, "User with this email already exists")
		}
		return nil, err
	}

	return users.PostUsers201JSONResponse{
		Id:    &createdUser.ID,
		Email: &createdUser.Email,
	}, nil
}

func (uh *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := uint(request.Id)

	if request.Body.Email == nil && request.Body.Password == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "No updates provided")
	}
	if request.Body.Email != nil && len(*request.Body.Email) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email cannot be empty")
	}
	if request.Body.Password != nil && len(*request.Body.Password) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Password cannot be empty")
	}

	updates := make(map[string]interface{})
	if request.Body.Email != nil {
		updates["email"] = *request.Body.Email
	}
	if request.Body.Password != nil {
		updates["password"] = *request.Body.Password
	}

	updatedUser, err := uh.Service.UpdateUserByID(id, updates)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return nil, err
	}

	return users.PatchUsersId200JSONResponse{
		Id:    &updatedUser.ID,
		Email: &updatedUser.Email,
	}, nil
}

func (uh *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := uint(request.Id)

	err := uh.Service.DeleteUserByID(id)
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}
