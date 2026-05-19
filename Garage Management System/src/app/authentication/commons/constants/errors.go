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

const (
	ErrPasswordFormat       = "Password must contain at least 8 characters long, at least one lower letter,at least one uppercase letter, at least one digit and at least one special character (@$!%*?&)."
	ErrConfirmPasswordMatch = "ConfirmPassword must match Password."
	ErrNewPasswordMatch     = "ConfirmPasword must match NewPassword"
)

const (
	ErrInvalidValue       = "invalid value for %s"
	ErrInvalidPanCard     = "Invalid PAN card format. It should be 5 uppercase letters, followed by 4 digits, and 1 uppercase letter."
	ErrInvalidPhoneNumber = "Phone number must be exactly 10 digits long and contain only numbers."
	ErrFieldRequired      = "%s is required."
	ErrInvalidEmail       = "Invalid value for Email"
)

const (
	ErrHashingPassword = "error hashing password: %w"
)

//header-token
const (
	InvalidTokenError            = "invalid token"
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

//logout
const (
	UserAlreadyLoggedoutError = "user already logged out"
	LogoutFailedError         = "logout failed"
)

// Database Constraint & Index Names
const (
	ErrUniqueConstraintViolation = "duplicate key value violates unique constraint"
	IndexCustomerssPanCard       = "idx_customers_pan_card"
	IndexCustomersEmail          = "idx_customers_email"
)

// Duplicate Entry Errors
const (
	ErrDuplicateEntry    = "already exists"
	ErrUsernameExists    = "username already exists"
	ErrUserAlreadyExists = "user already exists"
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

//Encrypt & Decrypt Erros
const (
	ErrFailedToEncrypt = "falied to encrpyt password"
)

//Signin and Token generation Errors
const (
	ErrInvalidEmailorPassword = "invalid email or password"
	ErrPasswordMismatch       = "password does not match %w"
	ErrAuthenticationFailed   = "authentication failed"
	ErrTokenGenerationFailed  = "failed to generate authentication tokens %s"
)

const (
	ErrSignInFailed      = "failed to sign in user"
	ErrUserNotFound      = "user not found"
	ErrIncorrectPassword = "entered password is not correct"
	ErrOtpsMismatch      = "OTPs did not match"
	ErrExpiredOtp        = "OTP expired"
	ErrIncorrectOtp      = "entered OTP is not correct"
	ErrInvalidOtp        = "OTP must be a 4 digit number"
	ErrJWTExpired        = "Token Expired"
	ErrJWTInvalid        = "Invalid Token"
	ErrJWTUnauthorized   = "Unauthorized"
)

const (
	ErrInvalidToken         = "invalid token"
	ErrPasswordChangeFailed = "failed to change password"
)

const (
	DatabaseQueryError = "database query error"
)
