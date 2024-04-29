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
	ErrInvalidAccount          = "user not found"
	ErrUserRoleNotExist        = "user role not exist"
	ErrPageNotFound            = "Page Not Found"
	ErrTokenNotExist           = "token not exist"
	ErrTokenNotFound           = "token not found"
	ErrInvalidToken            = "invalid token"
	ErrUnauthorizedAction      = "unauthorized action"
	ErrEmailAlreadyExist       = "email already exist"
	ErrNameAlreadyExist        = "Name Already Exist"
	ErrPasswordNotMatch        = "password not match"
	ErrConvertIdToInt          = "id must in number format"
	ErrPasswordContainUsername = "password cannot contain username"
	ErrInvalidEmail            = "the email is invalid"
	ErrInvalidTokenType        = "invalid token type"
	ErrTokenAlreadyExpired     = "token already expired"
	ErrRequiredOldPassword     = "old password is required"
	ErrInvalidNewPassword      = "new password is invalid"
	ErrOTPCodeNotFound         = "otp code is not found"

	ErrParseToFormatDate   = "please insert date with format : yyyy-mm-dd"
	ErrOrderStatusNotFound = "order status not found"

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
	ErrInvalidFullName        = "fullname can only contain characters and spaces"
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

	// review
	ErrReviewAlreadyCreated   = "you already create review for this transactions"
	ErrMaxUpload5File         = "upload file max 5 item"
	ErrNoReviewToProcess      = "no review data to process"
	ErrMustCompleteOrderFirst = "you must complete the order first"

	// wallet
	ErrWalletAlreadyRegistered = "wallet already registered"
	ErrPinAlreadyActivated     = "pin already activated"
	ErrPinWrong                = "wrong pin"
	ErrPinWrongBlocked         = "wrong pin, your wallet blocked"
	ErrWalletBlocked           = "your wallet blocked"
	ErrNotEnoughBalance        = "not enough balance"
	ErrPinLength               = "pin must filled by six number"
	ErrWalletNotActive         = "please activate your wallet first"
	ErrWalletNotRegistered     = "wallet not registered"
	ErrWalletNotFound          = "wallet is not found"

	// seller
	ErrSellerUnregistered                = "seller unregistered"
	ErrSellerAlreadyRegistered           = "seller already registered"
	ErrStoreNameMustNotBeEmpty           = "store name must not be empty"
	ErrDefaultStoreAddressMustNotBeEmpty = "default address must not be empty"
	ErrStoreNameAlreadyRegistered        = "store name already registered by other user"
	ErrSellerIdMustBeNumber              = "seller id must be number"
	ErrSellerNotFound                    = "seller is not found"

	ErrFieldNameEmpty     = "field name cannot be empty"
	ErrFieldEmailEmpty    = "field email cannot be empty"
	ErrFieldPasswordEmpty = "field password cannot be empty"

	ErrFieldDetailEmpty  = "field detail cannot be empty"
	ErrFieldCardNoEmpty  = "field card_no cannot be empty"
	ErrFieldUrbanIdEmpty = "field urban_id cannot be empty"

	// Address
	ErrAddressNotFound            = "data: address not exist"
	ErrCityNotFound               = "data: city not exist"
	ErrUrbanNotFound              = "data: urban not exist"
	ErrProvinceNotFound           = "data: province not exist"
	ErrDataAlreadyExist           = "data: address already exist"
	ErrCannotDeleteDefaultAddress = "cannot delete default address, please set it into another address first"

	// Product
	ErrMustInsertImage                    = "product image must filled"
	ErrVariantCountMustNumber             = "variant count must filled by number"
	ErrVariantMustNotNull                 = "you must fill variant"
	ErrVariantChoiceCountMustNumber       = "variant choice count must filled by number"
	ErrVariantIdMustNumber                = "err product variant id param must in number"
	ErrProductVariantCategoryIdMustNumber = "variant category id must filled by number"
	ErrStockDetailMustNumber              = "stock must filled by number"
	ErrStockCantFillWithNegativeNumber    = "stock must filled by positive number"
	ErrInvalidStockRequest                = "invalid stock request"
	ErrPriceDetailMustNumber              = "price must filled by number"
	ErrPriceMustBiggerThan0               = "price must bigger than 0"
	ErrNameVariant1NotFound               = "variant 1 name unsynchronized"
	ErrNameVariant2NotFound               = "variant 2 name unsynchronized"
	ErrDuplicateSku                       = "duplicate SKU"
	ErrDataNotYours                       = "data you delete not yours"
	ErrProductIdMustNumber                = "product id must number"
	ErrProductIdMustFilled                = "you must fill product id"
	ErrProductNotFromSeller               = "cant find product from seller"
	ErrProductNotFound                    = "product is not found"
	ErrWeightMin0AndMax30000              = "product weight must bigger than 0 and max 30.000"

	// Product detail
	ErrProductDetailMustFilled    = "product detail must filled"
	ErrProductDetailIdMustFilled  = "product detail id must filled"
	ErrProductDetailNotFound      = "product detail not found"
	ErrProductDetailQtyMustFilled = "product detail qty must filled"

	// Promotions
	ErrPromotionNameDuplicate          = "promotion name duplicate"
	ErrPromotionNotFound               = "promotion not found"
	ErrPromotionMustBiggerThan0        = "promotion must bigger than 0"
	ErrProductHaveMultiplePromotion    = "product already have promotion on that date"
	ErrActiveFromCantBiggerThanTo      = "active from cant bigger than active to"
	ErrActiveFromCantSmallerThanToday  = "active from cant smaller than today"
	ErrCantDeleteActiveOrDonePromotion = "cant delete active or done promotion"
	ErrPromotionTypeNotFound           = "promotion type not found"
	ErrPercentageMustBetween1Until100  = "discount with percentage must between 1 - 100"
	ErrDiscountMustLowerThanPrice      = "amount discount must lower than price"
	ErrDiscountMustBiggerThan0         = "amount discount must bigger than 0"
	ErrDeletePromotionOnProcess        = "can't delete on process promotion"
	ErrEditPromotionOnProcess          = "can't edit on process or done promotion"
	ErrDeletePromotionFinish           = "can't delete done promotion"
	ErrPromotionIdMustInNumber         = "promotion number must in number"

	// Voucher
	ErrVoucherTypeNotFound                   = "voucher type is not found"
	ErrVoucherNameRequired                   = "voucher name is required"
	ErrVoucherCodeRequired                   = "voucher code is required"
	ErrVoucherNotFound                       = "voucher is not found"
	ErrInvalidPercentageVoucherAmount        = "amount of voucher must be a minimal 0% and max 100%"
	ErrInvalidNominalVoucherAmount           = "amount of voucher must be greater or equal than 0"
	ErrActiveFromRequired                    = "a start period time is required"
	ErrActiveToRequired                      = "an end period of time is required"
	ErrInvalidActiveFrom                     = "a start period time is invalid"
	ErrInvalidActiveTo                       = "an end period of time is invalid"
	ErrInvalidPeriodTime                     = "period time is invalid"
	ErrInvalidMinAmount                      = "min amount is invalid"
	ErrInvalidMaxAmountDiscount              = "max amount discount is invalid"
	ErrNegativeQuota                         = "voucher quota must be greater or equal than 0"
	ErrCodeLenTooLong                        = "voucher code is too long"
	ErrVoucherAlreadyCreated                 = "voucher already created"
	ErrChangingVoucherCodeNotAllowed         = "changing voucher code is not allowed"
	ErrChangingMinAmountNotAllowed           = "changing voucher min-amount is only applicable in percentage voucher"
	ErrChangingMaxDiscountNotAllowed         = "changing voucher max-discount is only applicable in percentage voucher"
	ErrSellerVoucherCantUsedTwice            = "seller voucher cant used twice"
	ErrSellerVoucherEmpty                    = "seller voucher quota is empty, please faster next time"
	ErrAdminVoucherCantUsedTwice             = "marketplace voucher cant used twice"
	ErrAdminVoucherEmpty                     = "marketplace voucher quota is empty, please faster next time"
	ErrNotInSellerVoucherTimeRange           = "seller voucher cant used, not in voucher time range"
	ErrNotInAdminVoucherTimeRange            = "marketplace voucher cant used, not in voucher time range"
	ErrSellerVoucherMinAmountBiggerThanOrder = "seller voucher min amount must bigger than your total order"
	ErrAdminVoucherMinAmountBiggerThanOrder  = "marketplace voucher min amount must bigger than your total order"
	ErrActiveVoucherCannotEdited             = "active voucher cannot be edited"

	// Cart
	ErrCartQtyExceedProductQty        = "cart quantity exceed product quantity"
	ErrCannotBuyOwnProduct            = "cannot buy your own product"
	ErrPleaseChooseProductVariant     = "please choose the product variant"
	ErrPleaseCompleteYourAddressFirst = "please complete your addresss first"
	ErrCartItemQtyNotFound            = "cart item qty not found"

	// Favorite
	ErrFavoriteItemAlreadyAdded = "favorite item already added"
	ErrFavoriteItemNotFound     = "favorite item is not found"

	// Courier
	ErrCourierIdNotFound = "data: courier_id not exist"

	// Shipment Option
	ErrShipmentOptionIdNotFound = "data: shipment_option_id not exist"
	ErrShipmentOptionInvalid    = "invalid shipping option"

	// Shipment Fee
	ErrShipmentFeeNotFound = "shipment fee not found"

	// Payment
	ErrPaymentTypeInvalid                    = "payment type invalid"
	ErrNotEnoughStock                        = "not enough stock"
	ErrMustFillFullNameAndPhoneNumber        = "you must fill full name and phone number in your profile first"
	ErrDifferentStoreTotalPrice              = "different store total price"
	ErrDifferentProductPrice                 = "different product price"
	ErrDifferentTotalPrice                   = "different total price"
	ErrDifferentShipmentFee                  = "different shipment fee"
	ErrCantBuyFromYourOwnStore               = "cant buy from your own store"
	ErrDestinationAddressInvalid             = "destination address invalid"
	ErrOriginAddressInvalid                  = "origin address invalid"
	ErrPaymentIdMustInNumber                 = "payment id must in number"
	ErrOrderStatusIdMustInNumber             = "order status id must in number"
	ErrFilterByInvalid                       = "filter by invalid"
	ErrPaymentDataNotFound                   = "payment data not found"
	ErrInvoiceDataNotFound                   = "invoice data not found"
	ErrPaymentAlreadyUpdated                 = "payment already updated before"
	ErrInvalidOrderStatusRequest             = "invalid order status request"
	ErrCantCompleteOrderCauseOnprocessRefund = "cant complete order, there's still on process refund"
	ErrMustProcessOrder                      = "you must process the order first"
	ErrDescriptionLowerThan5Character        = "field description lower than 5 character"

	// Withdrawal
	ErrAmountMustBiggerThan0 = "amount must bigger than 0"

	// transaction
	ErrTransactionNotFound = "transaction not found"

	ErrGormDataUserIdNotExist       = "data: user_id not exist"
	ErrGormDataUserRoleIdNotExist   = "data: user_role_id not exist"
	ErrGormDataUserRoleNameNotFound = "data: role name is not found"
	ErrGormDataUserEmailNotExist    = "data: email not exist"
	ErrGormDataNameNotExist         = "data: name not exist"

	// Sealabs Pay
	ErrTransactionIdNotFound            = "transaction id not found"
	ErrFieldSealabsPayMustBe16Character = "field card_no must be exact 16 character"
	ErrSealabsPayAlreadyVerified        = "the sealabs pay already used by another user"
	ErrTransactionStatusNotInReady      = "transaction is not in ready status"
	ErrSignatureNotMatchWithActual      = "signature not match with the actual"
	ErrCannotUseAnotherPersonSealabsPay = "cannot use another person sealabs pay"
	ErrCardNoNotFound                   = "card no not found"
	ErrMinimumTopUpForSealabsPay1000    = "minimum top up for sealabs pay is 1000"
	ErrSealabsPayTransactionNotFound    = "sealabs pay transaction not found"
	ErrCannotConfirmOtherUserData       = "cannot confirm other user data"
	ErrCannotCancelOtherUserData        = "cannot cancel other user data"

	// rollback
	ErrRollbackNotFound = "rollback transaction not found"

	// Refund
	ErrPleseImportSupportingImage           = "please import the supporting image (min:1)"
	ErrImageCannotExceedFivePhoto           = "image cannot exceed 5 photo"
	ErrCannotRefundAnotherPersonData        = "cannot refund use another person data"
	ErrProductHaventOnReceivedStatus        = "the product haven't on received status"
	ErrPaymentTransactionAlreadyCompleted   = "the payment data already completed"
	ErrCannotRefundDueExceedOneWeek         = "refund cannot be proceed due to product already delivered longer than 1 Week"
	ErrCannotRefundDueAlreadyDeclinedOneDay = "refund cannot be proceed due to the previous refund already declined more than 1 day"
	ErrRefundAlreadyOnProcess               = "refund already on process"
	ErrTransactionAlreadyRefund             = "the transaction already refunded"
	ErrRefundIdCannotEmpty                  = "data: refund_id cannot be empty"
	ErrPaymentDetailIdCannotEmpty           = "data: payment_detail_id cannot be empty"
	ErrRefundCategoryIdCannotEmpty          = "data: refund_category_id cannot be empty"
	ErrReasonNotExceedTenCharacter          = "data: reason must be exceed ten character"
	ErrRefundDataNotFound                   = "refund data not found"
	ErrRefundCategoryDataNotFound           = "refund category data not found"
	ErrRefundDataNotInPendingStatus         = "current refund data not in waiting seller confirmation status"
	ErrRefundDataNotInApprovedStatus        = "current refund data not in approval status"
	ErrTxnIdNotFound                        = "data: txn_id not exist"

	// Product Category
	ErrProductCategoryDataNotFound                 = "product category data not found"
	ErrCannotCreateProductCategoryDueLevel3        = "cannot create product category due parent id already level 3"
	ErrSomethingWentWrongWhenCreateProductCategory = "something went wrong when create product category"
	ErrNotFoundLevel1ProductCategoryName           = "please fill product category level 1 name"
	ErrNotFoundLevel2ProductCategoryName           = "please fill product category level 2 name"
	ErrNotFoundLevel3ProductCategoryName           = "please fill product category level 3 name"
	ErrCannotHaveProductCategoryChildrenDueLevel3  = "cannot have product category children due cannot exceed 3 level"
	ErrParentIdCannotEmpty                         = "parent_id cannot empty"
	ErrLevel1ProductCategoryAlreadyExist           = "the level 1 product category already exist"
	ErrLevel2ProductCategoryAlreadyExist           = "the level 2 product category already exist"
	ErrLevel3ProductCategoryAlreadyExist           = "the level 3 product category already exist"
	ErrOnlyLevel1ProductCategoryCanAddImage        = "only level 1 product category can add image"
	ErrProductCategoryIdCannotBeEmpty              = "data product_category_id cannot be empty"
	ErrImageCannotExceedOnePhoto                   = "image cannot exceed 1 photo"
	ErrImageCannotBeEmpty                          = "image cannot be empty"
	ErrOnlyProductCategoryLevel1AddImage           = "only product category level 1 can add image"

	ErrPleaseRelogin = "your token expired, please relogin"
)

func (r *RespError) Error() string {
	return fmt.Sprintf("%d: %v", r.Code, r.Message)
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
