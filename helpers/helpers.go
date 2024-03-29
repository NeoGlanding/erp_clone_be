package helpers

import (
	"errors"
	"fmt"
	"time"

	"github.com/automa8e_clone/types"
	"github.com/go-playground/validator/v10"
)


func DestructValidationError(err *error) []types.ApiError {
	var returnValue []types.ApiError;
	if *err != nil {
        var ve validator.ValidationErrors
        if errors.As(*err, &ve) {
            out := make([]types.ApiError, len(ve))
            for i, fe := range ve {
                out[i] = types.ApiError{Field:fe.Field(), Message: msgForError(fe.Tag())}
            }
			returnValue = out
        }

    }
	return returnValue;
}

func msgForError(err string) string {
	switch err {
	case "required":
		return "This field cannot be empty"
	default:
		return "Invalid format"
	}
}

func FindTotalPage(totalRecords int64, pageSize int) int {
	result := int((totalRecords + int64(pageSize) - 1) / int64(pageSize))
	
	if result == 0 {
		return 1
	}

	return result
}

func StringArrayContains(list []string, element string) (bool) {
	for _, b := range list {
		if b == element {
			return true
		}
	}

	return false
}

func IsValidDateString(dateString string) bool {
	layout := "2006-01-02"
	_, err := time.Parse(layout, dateString)

	return err == nil
}

func FormatToTimestamps(dateString string) (time.Time, error) {
	timeString := "00:00:00.000000-07"

	// Combine the date and time strings
	dateTimeString := fmt.Sprintf("%s %s", dateString, timeString)

	// Define the layout for parsing
	layout := "2006-01-02 15:04:05.999999-07"

	// Parse the combined string into a time.Time value
	parsedTime, err := time.Parse(layout, dateTimeString)
	if err != nil {
		return parsedTime, err
	} else {
		return parsedTime, nil
	}
}