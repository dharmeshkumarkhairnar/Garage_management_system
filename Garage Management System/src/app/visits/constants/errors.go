package constants

const (
	VehicleCreationFailedError           = "vehicle creation failed"
	VehicleNumberPlateAlreadyExistsError = "vehicle with this number plate already exists"
	UserNotFoundError                    = "user not found"
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