package constants

const (
	VehicleCreationFailedError           = "vehicle creation failed"
	VehicleNumberPlateAlreadyExistsError = "vehicle with this number plate already exists"
	UserNotFoundError                    = "user not found"
	NumberPlateFormatError               = "Number Plate must be in `MH19BW3626` format"
	VehicleNotFoundError                 = "Vehicle Not Found"
	MechanicIDFormatError                = "Length of MechanicId must range between 1-32 followed by 3 digit numbers"
	ArrivalDateFormatError               = "Arrival Date must be in format YYYY-MM-DD"
	DeliveryDateFormatError              = "Delivery Date must be in format YYYY-MM-DD"
	DeliveryDateError                    = "Delivery Date must be greater than or equal to Arrival Date"
	ServicesLenghtError                  = "List of services cannot be empty"
)

// middleware errors
const (
	OperationFailed       = "operation failed"
	InvalidTokenError     = "token is invalid"
	MissingUserIDError    = "user id not found in token"
	MissingUserRoleError  = "user role not found token"
	RedisOperationError   = "error in redis operation"
	UserLoggedOutError    = "user is logged out"
	HeaderIsMissingError  = "header is missing"
	UnauthorizedUserError = "unauthorized user"
)
