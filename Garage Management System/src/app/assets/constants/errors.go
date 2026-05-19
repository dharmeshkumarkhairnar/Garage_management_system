package constants

const (
	MechanicAdditionFailedError = "mechanic addition failed"
	MechanicDeletionFailedError = "mechanic deletion failed"
	UserNotFoundError           = "user not found"
	DuplicateAadharNumberError  = "mechanic with this aadhar already exists"
)

// middleware errors
const (
	OperationFailed         = "operation failed"
	InvalidTokenError       = "token is invalid"
	MissingUserIDError      = "user id not found in token"
	MissingUserRoleError    = "user role not found token"
	RedisOperationError     = "error in redis operation"
	UserLoggedOutError      = "user is logged out"
	HeaderIsMissingError    = "header is missing"
	UnauthorizedUserError = "unauthorized user"
)
