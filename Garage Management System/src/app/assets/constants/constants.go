package constants

//Table related constants
const (
	MechanicsTableName = "mechanics"
)

//Success
const (
	MechanicAddedSuccessfully = "mechanic added successfully"
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

//logger messages
const(
	MechanicAddedInDB="mechanic added to DB"
)

