package response

// AppError represents a structured API error
type AppError struct {
	Status  int
	Message string
}

func (a AppError) Error() string {
	return a.Message
}

// Common error messages
const (
	ErrMsgInternalServerError                    = "Internal server error"
	ErrMsgBadRequest                             = "Bad request"
	ErrMsgUnauthorized                           = "Unauthorized"
	ErrMsgInvalidAccessToken                     = "Invalid access token"
	ErrMsgParamIdIsRequired                      = "ID parameter is required"
	YourRoleNotAllowedToAccessThisResource       = "Your role is not allowed to access this resource"
	YourPermissionNotAllowedToAccessThisResource = "Your permission is not allowed to access this resource"
)
