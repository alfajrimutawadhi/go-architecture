package helper

const (
	BadRequestMessage = "Bad Request"
	NotFoundMessage = "Not Found"
	InternalServerErrorMessage = "Internal Server Error"
	BadGatewayMessage = "Bad Gateway"
)

type BaseApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}