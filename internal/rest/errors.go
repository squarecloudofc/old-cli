package rest

import "fmt"

type SquareError string

const (
	AppNotFound   SquareError = "Application not found, verify your application ID and try again"
	AccessDenied  SquareError = "Your access is denied, make login with your api token using \"squarecloud login\" or verify if you have access for this action"
	UserNotFound  SquareError = "User not found, verify your user ID and try again"
	InvalidBuffer SquareError = "Unable to send buffer."
	InvalidFile   SquareError = "Unable to send the file"
	CommitError   SquareError = "Unable to commit to your application"
	DelayNow      SquareError = "You are in rate limit, try again later"
)

func ParseError(e *ApiResponse) (err string) {
	switch e.Code {
	case "APP_NOT_FOUND":
		err = string(AppNotFound)
	case "USER_NOT_FOUND":
		err = string(UserNotFound)
	case "ACCESS_DENIED":
		err = string(AccessDenied)
	case "INVALID_FILE":
		err = string(InvalidFile)
	case "INVALID_BUFFER":
		err = string(InvalidBuffer)
	case "COMMIT_ERROR":
		err = string(CommitError)
	case "DELAY_NOW":
		err = string(DelayNow)
	default:
		err = fmt.Sprintf("Square Cloud retorned error %s", e.Code)
	}

	return
}
