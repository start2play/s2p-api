package exceptions

import "errors"

var (
	INVALID_TOKEN             = errors.New("invalid_token")
	USER_NOT_EXISTS           = errors.New("user_not_exists")
	USER_NOT_AUTHORIZED       = errors.New("not_authorized")
	INVALID_EMAIL_OR_PASSWORD = errors.New("Email_or_Password_invalid")
)