package constants

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
