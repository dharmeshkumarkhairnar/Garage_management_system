package constants

//Table related constants
const (
	MechanicsTableName = "mechanics"
	AadharNumber       = "aadhar number"
)

//Success
const (
	MechanicAddedSuccessfully   = "mechanic added successfully"
	MechanicDeletedSuccessfully = "mechanic deleted successfully"
)

//database
const (
	AadharNumberCondition = "aadhar_number = ?"
)

//constant errors from Database
const (
	DuplicateAadharNumberDBError = "duplicate key value violates unique constraint \"idx_mechanics_aadhar_number\""
)

// Request Validation Errors
const (
	ErrInvalidPayload  = "invalid required payload"
	ErrUnexpectedValue = "unexpected value for the field."
)

//logger messages
const (
	MechanicAddedInDB   = "mechanic added to DB"
	MechanicDeletedInDB = "mechanic deleted to DB"
)

//middlware constants
const (
	User          = "user"
	Admin         = "admin"
	Header        = "header"
	Token         = "token"
	Subject       = "sub"
	UserRole      = "role"
	UserId        = "user_id"
	Authorization = "Authorization"
	Bearer        = "Bearer "
	ActiveToken   = "ACTIVE_TOKEN_%s"
	Redis         = "redis"
)
