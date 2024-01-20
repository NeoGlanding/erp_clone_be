package helpers

import (
	"errors"

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
	return int((totalRecords + int64(pageSize) - 1) / int64(pageSize))
}