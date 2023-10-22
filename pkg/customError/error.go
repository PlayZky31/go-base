package customError

import "fmt"

const (
	GeneralErrMessage      = "Oops! There's a hiccup on our server. Please wait a sec. 🛠️😅"
	BadRequestErrMessage   = "Oopsie! Something's off with your data. Can you check it again? 😊"
	DataNotFoundErrMessage = "Oops!! The data you're looking for isn't here. 🤷‍🔍"
)

type CustomError struct {
	Message  string            `json:"message"` // Human-readable message for clients
	Code     int               `json:"-"`       // HTTP Status code. We use `-` to skip json marshaling.
	Err      error             `json:"-"`
	ErrField map[string]string `json:"-"`
}

func (err CustomError) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}
	return err.Message
}

func ErrGeneral(err error) *CustomError {
	return &CustomError{
		Message: "",
		Code:    500,
		Err:     err,
	}
}

func ErrQuery(err error) *CustomError {
	return &CustomError{
		Message: GeneralErrMessage,
		Code:    500,
		Err:     fmt.Errorf("query error: %s", err.Error()),
	}
}

func ErrNotFound(title string) *CustomError {
	return &CustomError{
		Message: DataNotFoundErrMessage,
		Code:    404,
		Err:     fmt.Errorf("data not found: %s", title),
	}
}

func ErrBadRequestFields(errField map[string]string) *CustomError {
	return &CustomError{
		Message:  BadRequestErrMessage,
		Code:     400,
		Err:      fmt.Errorf("bad request"),
		ErrField: errField,
	}
}

func ErrBadRequest(err error) *CustomError {
	return &CustomError{
		Message: BadRequestErrMessage,
		Code:    400,
		Err:     fmt.Errorf("bad request: %s", err.Error()),
	}
}

func ErrBadRequestWithCustomMsg(msg string) *CustomError {
	return &CustomError{
		Message: msg,
		Code:    400,
		Err:     fmt.Errorf("bad request: %s", msg),
	}
}
