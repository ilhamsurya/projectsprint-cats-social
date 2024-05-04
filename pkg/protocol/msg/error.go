package msg

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type RespError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

const (
	ErrInvalidRequest          = "invalid request"
	ErrInvalidJsonName         = "invalid json name"
	ErrUserNotFound            = "User is not found"
	ErrUserRoleNotExist        = "user role not exist"
	ErrPageNotFound            = "Page Not Found"
	ErrTokenNotExist           = "token not exist"
	ErrTokenNotFound           = "token not found"
	ErrInvalidToken            = "invalid token"
	ErrUnauthorizedAction      = "unauthorized action"
	ErrEmailAlreadyExist       = "email already exist"
	ErrNameAlreadyExist        = "Name Already Exist"
	ErrWrongPassword           = "Password is wrong"
	ErrConvertIdToInt          = "id must in number format"
	ErrPasswordContainUsername = "password cannot contain username"
	ErrInvalidEmail            = "the email is invalid"
	ErrInvalidTokenType        = "invalid token type"
	ErrTokenAlreadyExpired     = "token already expired"
	ErrRequiredOldPassword     = "old password is required"
	ErrInvalidPassword         = "Password should contains characters and must be between 5 and 15 characters long"
	ErrOTPCodeNotFound         = "otp code is not found"

	ErrParseToFormatDate    = "please insert date with format : yyyy-mm-dd"
	ErrInvalidSigningMethod = "Signing method invalid"
	// filter
	ErrLimitNotNumber              = "limit must filled by number"
	ErrLimitMustBetween0Until100   = "limit must filled between 1-100"
	ErrInvalidSortRequest          = "sort request not found"
	ErrInvalidLevelRequest         = "level request invalid"
	ErrInvalidFilterRequest        = "filter request invalid"
	ErrPageNotNumber               = "page must filled by number"
	ErrStarNotNumber               = "star must filled by number"
	ErrIdSellerNotNumber           = "seller id must filled by number"
	ErrIdCityNotNumber             = "city id must filled by number"
	ErrIdProvinceNotNumber         = "province id must filled by number"
	ErrMinPriceNotNumber           = "min price must filled by number"
	ErrMaxPriceNotNumber           = "max price must filled by number"
	ErrIdCategoryNotNumber         = "id category must filled by number"
	ErrMinRatingNotNumber          = "rating must filled by number"
	ErrRatingMustBetween1Until5    = "rating must filled between 1 until 5"
	ErrMinPriceMinimal0            = "min price cant fill with minus"
	ErrMaxPriceMinimal1            = "max price must bigger than 0"
	ErrMinRatingMustBetween0Until5 = "min rating must between 0 - 5"
	ErrStatusNotFound              = "status not found"

	// user
	ErrUserAlreadyExist       = "user already exists"
	ErrUserNotExist           = "user not exists"
	ErrUsernameAlreadyExist   = "user name already exists"
	ErrUserNotASeller         = "user not a seller"
	ErrRequiredFullName       = "fullname is required"
	ErrInvalidFullName        = "name can only contain characters and spaces, and must be between 5 and 15 characters long"
	ErrFullNameAlreadyUsed    = "fullname already used"
	ErrRequiredPhoneNumber    = "phone number is required"
	ErrInvalidPhoneNumber     = "phone number is invalid"
	ErrPhoneNumberAlreadyUsed = "phone number already used"
	ErrInvalidBirthDate       = "invalid birth date"
	ErrRequiredBirthDate      = "birth date is required"
	ErrRequiredGender         = "gender is required"
	ErrInvalidGenderFormat    = "gender format is invalid"
	ErrInvalidFormatFile      = "uploaded file must be in image format"
	ErrUnsupportedImgFormat   = "image format is unsupported"
	ErrProfileRecordNotFound  = "profile is not found"
	ErrRequiredUsername       = "username cannot be empty"
	ErrWeakPassword           = "password should contain alphanumeric, symbol and one character to be uppercase with min length is 6"
	ErrOldPassword            = "old password cannot be used"

	ErrPleaseRelogin = "your token expired, please relogin"
)

func (r *RespError) Error() string {
	return fmt.Sprintf("%d: %v", r.Code, r.Message)
}

func UnwrapRespError(err interface{}) RespError {
	respError, ok := err.(*RespError)
	if !ok {
		panic("fail when casting the respError type")
	}

	return *respError
}

func InternalServerError(msg string) error {
	log.Printf(msg)

	return &RespError{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}
}

func BadRequest(msg string) error {
	return &RespError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func NotFound(msg string) error {
	return &RespError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func Unauthorization(msg string) error {
	return &RespError{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}
