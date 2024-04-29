package msg

const (
	GetAllResponse                      = "get all data successfully"
	GetByIDResponse                     = "Get Data By ID Successfully"
	GetByCodeResponse                   = "Get Data By Code Successfully"
	GetDataResponse                     = "Get Data Successfully"
	CreateResponse                      = "Data Created Successfully"
	UpdateResponse                      = "Data Updated Successfully"
	DeleteResponse                      = "Data Deleted Successfully"
	ConfirmResponse                     = "Data Confirmed Successfully"
	LoginResponse                       = "Login Successfully"
	SignInWithGoogle                    = "signin with google successfully"
	RegisterResponse                    = "Register Successfully"
	UpdateUserProfileResponse           = "User Profile Successfully Updated"
	AccessTokenSuccessfully             = "Access Token Successfully Generated"
	ResetPasswordSuccessfully           = "reset password Successfully"
	ForgotPasswordRequestedSuccessfully = "forgot password requested successfully"
	RequestChangePasswordSuccessfully   = "request change password successfully"
	PasswordSuccessfullyChanged         = "password successfully changed"
	WithdrawalSuccess                   = "withdrawal success"
	// seller
	GetByIdUserResponse             = "Get Data By ID User Successfully"
	SetAsSellerResponse             = "Set As Seller Successfully"
	VerifySeaLabsPay                = "please verify sealabs pay in this url"
	SealabsPaySuccessfullyVerified  = "sealabs pay successfully verified"
	SealabsPaySuccessfullyConfirmed = "sealabs pay successfully confirmed"
	SealabsPaySuccessfullyCancelled = "sealabs pay successfully cancelled"
	SealabsPaySuccessfullyTopUp     = "top up with sealabs pay successfully"
	RefundApproveResponse           = "Refund Approved Successfully"
	RefundDeclinedResponse          = "Refund Declined Successfully"
	RefundSuccessfully              = "Transaction Refunded Successfully"
	UserWalletSuccessfullyVerified  = "user wallet successfully verified"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

func ReturnResult(message string, data interface{}) Response {
	res := Response{
		Message: message,
		Data:    data,
	}
	return res
}
