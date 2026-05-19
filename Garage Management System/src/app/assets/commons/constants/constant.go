package constants

//Database related
const (
	DSNString = "host=%s port=%s dbname=%s user=%s password=%s TimeZone=%s"
	AccessKey = "Apr/meTe4sxpBwxb36ISTRNnHc4y+Y34KjQ/ntwB1Kw="
	// PasswordRegex = "^(?=.[a-z])(?=.[A-Z])(?=.\\d)(?=.[\\W_]).+$"
)

//Authentications API URL Keys
const (
	ServiceName       = "authentication"
	PortDefaultValude = 8081
)

const (
	ServiceAddedSuccessfully   = "service added successfully"
	ServiceDeletedSuccessfully = "service deleted successfully"
)

// Database table name & field names for users
const (
	UsersTableName = "users"
	Fieldemail     = "email"
	Username       = "username = ?"
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
	Origin      = "Origin"
	ContentType = "Content-type"
	// Authorization = "Authorization"
	// Bearer        = "Bearer"
	// Token         = "token"
	// Server        = "server"
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
	Server  = "server"
	Service = "service"
)

//Table related constants
const (
	MechanicsTableName        = "mechanics"
	FieldMechanicAadharNumber = "AadharNumber"
	FieldMechanicPhoneNumber  = "PhoneNumber"
	FieldMechanicName         = "Name"
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
