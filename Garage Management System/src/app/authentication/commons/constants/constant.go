package constants

//Database related
const (
	DSNString = "host=%s port=%s dbname=%s user=%s password=%s TimeZone=%s"
	AccessKey = "Apr/meTe4sxpBwxb36ISTRNnHc4y+Y34KjQ/ntwB1Kw="
)

//Authentications API URL Keys
const (
	ServiceName       = "authentication"
	PortDefaultValude = 8081
)

// Database table name & field names for Customers
const (
	CustomersTableName = "customers"
	CustomerEmail      = "email"
	CustomerName       = "name = ?"
	UsersTableName   = "users"
	Fieldemail       = "email"
	Username         = "username = ?"
	EmailPlaceholder = "email = ?"
)

// Success message for user
const (
	UserCreationSuccessMsg     = "Customer created successfully"
	CustomerLoggedInSuccessMsg = "Customer logged in successfully"
	OtpValidatedSuccessMsg     = "OTP validated successfully"
	PasswordChangedSuccessMsg  = "Password changed succefully"
	LogoutSuccessfulMsg        = "Customer logged out successfully"
)

//Swagger Titile
const SwaggerTitle = "Stock Broker Application API"

const EmailorPasswordField = "email_or_password"

//Cookies
const (
	Name     = "refresh_token"
	Time     = 30 * 24 * 60 * 60
	Path     = "/"
	Domain   = ""
	Secure   = true
	HttpOnly = true
)

const (
	RunningServerPort = "Running Server on port : %v"
)

const (
	PasswordRegex    = `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	EmailRegex       = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z]{2,}$`
	UppercaseRegex   = `[A-Z]`
	DigitRegex       = `\d`
	SpecialCharRegex = `[@$!%*?&]`
	LowercaseRegex   = `[a-z]`
	OtpRegexp        = "^[0-9]{4}$"
)

//customer
const (
	FieldPassword            = "Password"
	FieldConfirmPassword     = "ConfirmPassword"
	FieldStrongPassword      = "strongPassword"
	FieldCustomerPhoneNumber = "PhoneNumber"
	FieldCustomerEmail       = "Email"
	FieldCustomerName        = "Name"
)

// Migration success Message
const (
	MsgDBMigrationSuccess = "Database migration completed successfully!"
)

const (
	Postgres = "postgres"
	JWT      = "jwt"
	Yaml     = "yaml"
)

// Database Keys
const (
	Interal = "internal"
)

// Paths
const (
	BaseConfig = "../../config"
	RootConfig = "./src/config"
)

// Origin
const (
	AllowedOrigin = "*"
)

// Method
const (
	POST = "POST"
	GET  = "GET"
)

// Header
const (
	Origin        = "Origin"
	ContentType   = "Content-type"
	Authorization = "Authorization"
	Bearer        = "Bearer"
	Token         = "token"
	Server        = "server"
)

// redis
const (
	RedisAddr     = "127.0.0.1:6379"
	RedisPassword = ""
	RedisDB       = 0
)

// claims
const (
	Exp        = "exp"
	Sub        = "sub"
	JTI        = "jti"
	DeviceType = "device_type"
)

const (
	RedisTokenCacheKey = "ACTIVE_TOKEN_%s"
)
