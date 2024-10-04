package customError

import "fmt"

type Http struct {
	Description string      `json:"description,omitempty"`
	StatusCode  int         `json:"statusCode"`
	Response    interface{} `json:"response"`
}

type Custom struct {
	Error error `json:"error,omitempty"`
}

func (e Http) Error() string {
	return fmt.Sprintf("description: %s,	statusCode: %d", e.Description, e.StatusCode)
}

func HttpError(description string, statusCode int) Http {
	return Http{
		Description: description,
		StatusCode:  statusCode,
	}
}

func HttpErrorResponse(description string, statusCode int, response interface{}) Http {
	return Http{
		Description: description,
		StatusCode:  statusCode,
		Response:    response,
	}
}
