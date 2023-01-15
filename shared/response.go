package shared

type ErrorResponse struct {
	Code int
	Err  string
} //@name ErrorResponse

type LoginSuccessResponse struct {
	Code    int
	Message string
	Token   string
} //@name LoginSuccessResponse
