package constants

// Database Transaction & Connection Errors
const (
	ErrBeginTx            = "failed to begin database transaction: %w"
	ErrCommitTx           = "failed to commit database transaction: %w"
	ErrDBConnectionFailed = "Error connecting to database: %s"
	ErrInternalServer     = "internal server error"
)

// Database Initialization & Config Errors
const (
	ErrDBInitFailed             = "Error initializing database: %s"
	ErrDBMigrationFailed        = "Error migrating database: %s"
	ErrLoadConfigFailed         = "failed to load config: %v"
	ErrPostgresConnectionFailed = "failed to connect with the postgres: %s"
	ErrReadConfigFailed         = "failed to read the config file: %s"
	ErrUnmarshallConfigFailed   = "failed to unmarshal the config file %s"
	ErrJWTConfigReadFailed      = "failed to read the JWT config file %s"
	ErrRedisInitFailed          = "failed to initialize redis %s"
)

//header-token
const (
	// InvalidTokenError            = "invalid token"
	AuthHeaderMissingError       = "missing authorization header"
	InvalidAuthFormatError       = "invalid authorization format"
	TokenAlreadyBlacklistedError = "token is already blacklisted"
	UnexpectedSigningMethod      = "unexpected signing method"
	InvalidExpClaimsError        = "invalid exp claims"
	InvalidSubClaimsError        = "invalid sub claims"
	InvlalidJTIClaimsError       = "invalid jti claims"
	InvalidDeviceTypeClaimsError = "invalid deviceType claims"
	TokenExpiredError            = "token is already expired"
	RedisClientNotInitialized    = "redis client not initialized"
	SessionExpiredError          = "session expired or logged out"
	SessionSuspendedError        = "session suspended due to new login"
	NewLoginDetectedError        = "new login detected on another device"
)

// Database Constraint & Index Names
const (
	ErrUniqueConstraintViolation = "duplicate key value violates unique constraint"
	IndexServiceName             = "idx_service_masters_service"
)

// General Errors
const (
	ErrConflict           = "conflict"
	ErrUserCreationFailed = "failed to create user"
)

// Request Validation Errors
const (
	ErrInvalidPayload  = "invalid required payload"
	ErrUnexpectedValue = "unexpected value for the field."
)

const (
	ErrJWTExpired      = "Token Expired"
	ErrJWTInvalid      = "Invalid Token"
	ErrJWTUnauthorized = "Unauthorized"
)

const (
	DatabaseQueryError = "database query error"
)

// Duplicate Entry Errors
const (
	ErrDuplicateEntry       = "already exists"
	ErrServiceAlreadyExists = "service already exists"
)

const (
	ServiceDoesNotExist   = "service does not exist"
	FailedToDeleteService = "failed to delete service"
)

const (
	MechanicAdditionFailedError = "mechanic addition failed"
	MechanicDeletionFailedError = "mechanic deletion failed"
	UserNotFoundError           = "user not found"
	DuplicateAadharNumberError  = "mechanic with this aadhar already exists"
)

// middleware errors
const (
	OperationFailed       = "operation failed"
	InvalidTokenError     = "token is invalid"
	MissingUserIDError    = "user id not found in token"
	MissingUserRoleError  = "user role not found in token"
	RedisOperationError   = "error in redis operation"
	UserLoggedOutError    = "user is logged out"
	HeaderIsMissingError  = "header is missing"
	UnauthorizedUserError = "unauthorized user"
)
