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
	ServiceAddedSuccessfully = "service added successfully"
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
	Origin        = "Origin"
	ContentType   = "Content-type"
	Authorization = "Authorization"
	Bearer        = "Bearer"
	Token         = "token"
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
