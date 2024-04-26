package errors

type ApplicationError struct {
	Message string `json:"message"`
	Cause   error  `json:"cause"`
}

func (e ApplicationError) Error() string {
	return e.Message
}

func NewApplicationError(message string, cause error) ApplicationError {
	return ApplicationError{
		Message: message,
		Cause:   cause,
	}
}

type UnprocessableError struct {
	ApplicationError
}

func NewUnprocessableError(message string) UnprocessableError {
	return UnprocessableError{
		ApplicationError: ApplicationError{
			Message: message,
		},
	}
}

type BadRequestError struct {
	ApplicationError
}

func NewBadRequestError(message string, cause error) BadRequestError {
	return BadRequestError{
		ApplicationError: ApplicationError{
			Message: message,
			Cause:   cause,
		},
	}
}

type NotFoundError struct {
	ApplicationError
}

func NewNotFoundError(message string, cause error) NotFoundError {
	return NotFoundError{
		ApplicationError: ApplicationError{
			Message: message,
			Cause:   cause,
		},
	}
}
