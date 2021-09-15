package common

var (
	/*
		common errors
	*/
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrToken            = &Errno{Code: 10003, Message: "Invalid Token."}
	ErrTokenExpired     = &Errno{Code: 10004, Message: "Token is expired."}
	ErrTokenNotValidYet = &Errno{Code: 10005, Message: "Token not active yet."}
	ErrTokenMalformed   = &Errno{Code: 10006, Message: "That's not even a token."}
	ErrTokenInvalid     = &Errno{Code: 10007, Message: "Couldn't handle this token:"}
	ErrGenerateToken    = &Errno{Code: 10008, Message: "Generate Token is wrong."}

	/*
		system errors
	*/
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrEncryption = &Errno{Code: 20003, Message: "encryption error."}
	ErrLoadPolicy = &Errno{Code: 20004, Message: "load policy error."}

	/*
		person errors
	*/
	ErrUserNotFound         = &Errno{Code: 20101, Message: "The user was not found."}
	ErrPasswordIncorrect    = &Errno{Code: 20102, Message: "The password was incorrect."}
	ErrUserExisted          = &Errno{Code: 20103, Message: "The user was existed."}
	ErrArticleNotExisted    = &Errno{Code: 20104, Message: "The Article was not existed."}
	ErrRoleExisted          = &Errno{Code: 20105, Message: "The role was existed."}
	ErrRoleNotExisted       = &Errno{Code: 20106, Message: "The role was not existed."}
	ErrPermissionNotExisted = &Errno{Code: 20107, Message: "The Permission was not existed."}
	ErrAccessDenied         = &Errno{Code: 20108, Message: "Access denied."}
	ErrUserExistedInRole    = &Errno{Code: 20109, Message: "The user already exists in the role."}
	ErrRelationshipNotExisted = &Errno{Code: 20110, Message: "The relationship does not exist."}
)
