package validations

import (
	"fmt"
	Assestconst "garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/app/authentication/commons/constants"
	VisitConst "garage_management_system/src/app/visits/commons/constants"
	ReqModel "garage_management_system/src/app/visits/models"
	"garage_management_system/src/models"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
)

var bffValidator *validator.Validate

func FormatValidationErrors(err error) ([]models.ErrorMessage, string) {
	var validationErrors []models.ErrorMessage
	var validationErrorsStr string

	for _, err := range err.(validator.ValidationErrors) {
		var errorMsg string
		fieldName := err.Field()
		if err.Tag() == "required" {
			fieldName = strings.ToLower(fieldName)
			errorMsg = fmt.Sprintf(constants.ErrFieldRequired, fieldName)
		} else {
			switch err.Field() {
			case constants.FieldPassword:
				errorMsg = constants.ErrPasswordFormat
			case constants.FieldConfirmPassword:
				if err.Tag() == "eqfield" {
					errorMsg = constants.ErrConfirmPasswordMatch
				}
			case constants.FieldCustomerEmail:
				if err.Tag() == "required" {
					errorMsg = constants.ErrInvalidEmail
				}
			case Assestconst.FieldMechanicAadharNumber:
				errorMsg = Assestconst.AadharFormatError
			case VisitConst.FieldServices:
				errorMsg = VisitConst.ServicesLenghtError
			case VisitConst.FieldNumberPlate:
				errorMsg = VisitConst.NumberPlateFormatError
			case VisitConst.FieldMechanicID:
				errorMsg = VisitConst.MechanicIDFormatError
			case VisitConst.FieldArrivalDate:
				errorMsg = VisitConst.ArrivalDateFormatError
			case VisitConst.FieldDeliveryDate:
				if err.Tag() == "gtefield" {
					errorMsg = VisitConst.DeliveryDateError
				} else {
					errorMsg = VisitConst.DeliveryDateFormatError
				}
			default:
				errorMsg = fmt.Sprintf(constants.ErrInvalidValue, err.Field())
			}
		}

		validationErrors = append(validationErrors, models.ErrorMessage{
			Key:          fieldName,
			ErrorMessage: errorMsg,
		})
		validationErrorsStr += fieldName + " is invalid; "
	}

	return validationErrors, validationErrorsStr
}

func strongPasswordValidator(f1 validator.FieldLevel) bool {
	re := regexp2.MustCompile(constants.PasswordRegex, 0)
	matched, _ := re.MatchString(f1.Field().String())
	return matched
}

func aadharFormatValidator(f1 validator.FieldLevel) bool {
	aadhar := f1.Field().String()
	if len(aadhar) != 12 {
		return false
	}
	_, err := strconv.Atoi(aadhar)
	if err != nil {
		return false
	}
	return true
}

func serviceLengthValidator(f1 validator.FieldLevel) bool {
	services := f1.Field().Interface().([]string)
	if len(services) == 0 {
		return false
	}
	return true
}

func deliveryDateFormatValidator(s1 validator.StructLevel) {
	request, ok := s1.Current().Interface().(ReqModel.BFFAddVisitRecordsRequest)
	if !ok {
		return
	}

	arrivalTime, err := time.Parse(time.DateOnly, request.ArrivalDate)
	if err != nil {
		return
	}

	deliveryTime, err := time.Parse(time.DateOnly, request.DeliveryDate)
	if err != nil {
		return
	}

	if deliveryTime.Before(arrivalTime) {
		s1.ReportError(
			reflect.ValueOf(request.DeliveryDate),
			"DeliveryDate",
			"delivery_date",
			"gtefield",
			"ArrivalDate",
		)
	}
}

func mechanicIDFormatValidator(f1 validator.FieldLevel) bool {
	re := regexp2.MustCompile(VisitConst.MechacnicIdRegex, 0)
	matched, _ := re.MatchString(f1.Field().String())
	return matched
}

func numberPlateFormatValidator(f1 validator.FieldLevel) bool {
	re := regexp2.MustCompile(VisitConst.NumberPlateRegex, 0)
	matched, _ := re.MatchString(f1.Field().String())
	return matched
}

func IsEmailValid(f1 validator.FieldLevel) bool {
	email := f1.Field().String()
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	domainParts := strings.Split(parts[1], ".")

	if len(domainParts) < 2 || len(domainParts) > 3 {
		return false
	}

	for i := 0; i < len(domainParts)-1; i++ {
		for j := i + 1; j < len(domainParts); j++ {
			if domainParts[i] == domainParts[j] {
				return false
			}
		}
	}

	EmailRegex := regexp.MustCompile(constants.EmailRegex)

	return EmailRegex.MatchString(email)
}

func init() {
	bffValidator = validator.New()
	bffValidator.RegisterValidation("strongPassword", strongPasswordValidator)
	bffValidator.RegisterValidation("Email", IsEmailValid)
	bffValidator.RegisterValidation("aadharformat", aadharFormatValidator)
	bffValidator.RegisterValidation("numPlateFormat", numberPlateFormatValidator)
	bffValidator.RegisterValidation("mechIDFormat", mechanicIDFormatValidator)
	bffValidator.RegisterValidation("serviceLength", serviceLengthValidator)
	bffValidator.RegisterStructValidation(deliveryDateFormatValidator, ReqModel.BFFAddVisitRecordsRequest{})
}

func GetBFFValidator() *validator.Validate {
	return bffValidator
}
