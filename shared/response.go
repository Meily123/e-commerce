package shared

import "WebAPI/model"

type ErrorResponse struct {
	Code  int
	Error string
} //@name ErrorResponse

type LoginSuccessResponse struct {
	Code    int
	Message string
	Token   string
} //@name LoginSuccessResponse

type SuccessUserResponse struct {
	code    int
	massage string
	user    model.UserResponse
} //@name SuccessUserResponse

type SuccessListUserResponse struct {
	code    int
	massage string
	user    []model.UserResponse
} //@name SuccessListUserResponse

type SuccessResponse struct {
	Code    int         `json:"code" `
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
} //@name SuccessResponse

func ListUserRenderToResponse(users []model.User) []model.UserResponse {
	var renderedList []model.UserResponse
	for _, user := range users {
		renderedList = append(renderedList, UserRenderToResponse(user))
	}
	return renderedList
}

func UserRenderToResponse(user model.User) model.UserResponse {
	return model.UserResponse{
		Id:       user.Id.String(),
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
	}
}
