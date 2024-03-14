package common

type errorMessageType struct {
	FAILED                 string
	ERROR                  string
	FORBIDDEN_RESOURCE     string
	UNAUTHORIZED           string
	UNABLE_ACCESS_RESOURCE string
}

var ErrorMessage errorMessageType = errorMessageType{
	FAILED:                 "Process failed",
	ERROR:                  "Process error",
	FORBIDDEN_RESOURCE:     "Forbidden Resource",
	UNAUTHORIZED:           "Unauthorized Access",
	UNABLE_ACCESS_RESOURCE: "Unable to access resource",
}
