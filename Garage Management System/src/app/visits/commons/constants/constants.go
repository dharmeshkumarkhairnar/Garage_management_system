package constants

//Table related constants
const (
	VehiclesTableName      = "vehicles"
	VisitRecordTableName   = "visit_records"
	VisitServiceTableName  = "visit_services"
	ServiceMasterTableName = "service_masters"
	MechanicsTableName     = "mechanics"
	VisitRecordID          = "visit_record_id"
	ServiceMasterID        = "service_master_id"
)

//Success
const (
	VehicleCreationSuccess     = "vehicle created successfully"
	BillGenerationSuccess      = "Bill Generated Successfully"
	VisitRecordCreationSuccess = "visit record created successfully"
)

//constant errors from Database
const (
	DuplicateNumberPlateError = "duplicate key value violates unique constraint \"idx_vehicles_number_plate\""
	CustomerNotFoundError     = "insert or update on table \"vehicles\" violates foreign key constraint \"fk_vehicles_cust_id\""
)

// Request Validation Errors
const (
	ErrInvalidPayload  = "invalid required payload"
	ErrUnexpectedValue = "unexpected value for the field."
)

//constants
const (
	User                 = "user"
	NumberPlate          = "number_plate = ?"
	Vehicle              = "vehicle"
	Mechanic             = "mechanic"
	Services             = "services"
	MechIDCondition      = "mech_id = ?"
	ServiceIDINCondition = "service IN ?"
)

//middlware constants
const (
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

//validation
const (
	NumberPlateRegex = "^[A-Z]{2}[0-9]{2}[A-Z]{2}[0-9]{4}$"
	MechacnicIdRegex = "^[a-z]{1,32}[0-9]{3}$"
)

//database
const (
	FieldNumberPlate   = "NumberPlate"
	FieldMechanicID    = "MechanicId"
	FieldArrivalDate   = "ArrivalDate"
	FieldDeliveryDate  = "DeliveryDate"
	FieldServices      = "Services"
	FieldVisitRecordID = "visit_record_id"
)
