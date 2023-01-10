package myerr

//!错误码定义
var (
	//*common errors
	OK                  = &ErrNum{Code: 0, Message: "OK"}
	InternalServerError = &ErrNum{Code: 30001, Message: "Internal server error"}       //?内部错误
	ErrBind             = &ErrNum{Code: 30002, Message: "Can't convert to the struct"} //无法转换结构体
	ErrDatabase         = &ErrNum{Code: 30002, Message: "Database error"}
	ErrValidation       = &ErrNum{Code: 30001, Message: "Validation failed"}
	ErrEncrypt          = &ErrNum{Code: 30101, Message: "Error occurred while ecrypting the user password"}

	//?uer errors
	ErrAccountNotFound = &ErrNum{Code: 50102, Message: "User not found"}
	ErrPassword        = &ErrNum{Code: 50103, Message: "Password error"}
	ErrAccountEmpty    = &ErrNum{Code: 50104, Message: "Username can't be empty"}
	ErrPasswordEmpty   = &ErrNum{Code: 50103, Message: "Password can't be empty"}
	ErrMissingHeader   = &ErrNum{Code: 50104, Message: "Header not found"}
	ErrToken           = &ErrNum{Code: 50105, Message: "Token error"}
	PassParamCheck     = &ErrNum{Code: 60000, Message: "Parameter validtion passed"}
)
